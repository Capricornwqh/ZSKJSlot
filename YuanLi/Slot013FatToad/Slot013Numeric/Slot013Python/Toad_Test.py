import multiprocessing  # 用于多进程测试，提高测试效率
import datetime  # 用于记录测试时间和创建时间戳目录
import json  # 用于数据序列化和反序列化
import copy  # 用于深拷贝对象
import os  # 用于文件和目录操作
import random  # 用于生成随机数，模拟游戏随机性

import numpy  # 用于计算方差等统计数据

import Toad_Slot as Game_Slot  # 导入蟾蜍老虎机游戏核心逻辑
import Toad_Config as Config  # 导入游戏配置参数
import Slot  # 导入基础老虎机功能
import Const as Const  # 导入常量定义
import base_free_trigger_data  # 导入测试数据
import prettytable as pt  # 用于格式化输出表格

# 统计数据字典：用于记录测试过程中的各项指标
data = {
    Const.S_Test_Time: 0,  # 测试总次数

    Const.S_Bet: 0,  # 总下注额
    Const.S_Win: 0,  # 总赢得金额

    Const.S_Base_Win: 0,  # 基础游戏赢得金额
    Const.S_SC_Win: 0,  # Scatter赢得金额

    Const.S_Base_Hit: 0,  # 基础游戏中奖次数
    Const.S_Base_Sym_Win: {  # 各符号中奖统计（按符号类型和连线数）
        Config.H1: [0, 0, 0, 0, 0, 0],  # H1高分符号的3连、4连、5连等中奖次数
        Config.H2: [0, 0, 0, 0, 0, 0],  # H2高分符号
        Config.H3: [0, 0, 0, 0, 0, 0],  # H3高分符号
        Config.H4: [0, 0, 0, 0, 0, 0],  # H4高分符号
        Config.LA: [0, 0, 0, 0, 0, 0],  # A符号的中奖统计
        Config.LK: [0, 0, 0, 0, 0, 0],  # K符号的中奖统计
        Config.LQ: [0, 0, 0, 0, 0, 0],  # Q符号的中奖统计
        Config.LJ: [0, 0, 0, 0, 0, 0],  # J符号的中奖统计
    },

    Const.S_Free_Hit: 0,  # 触发免费游戏次数
    Const.S_Free_Win_Hit: 0,  # 免费游戏中有中奖的次数
    Const.S_Free_Win: 0,  # 免费游戏总赢得金额
    Const.S_Free_Spin: 0,  # 免费游戏总旋转次数


    Const.S_Feature_Hit: 0,  # 特殊功能触发次数
    'Max_Win': [0 for _ in range(40)],  # 各进程记录的最大赢额倍数
    'Max_Win_Times': 0,  # 达到限制最大赢额的次数
    'Variance': [0 for _ in range(40)],  # 各进程的方差值
    
    # 不同赢额倍数区间的统计 [总体, 免费游戏, 基础野生增益]
    'win_0': [0, 0, 0],  # 0倍赌注的次数
    'win_0_1': [0, 0, 0],  # 0-1倍赌注的次数
    'win_1_2': [0, 0, 0],  # 1-2倍赌注的次数
    'win_2_3': [0, 0, 0],  # 2-3倍赌注的次数
    'win_3_4': [0, 0, 0],  # 3-4倍赌注的次数
    'win_4_5': [0, 0, 0],  # 4-5倍赌注的次数
    'win_5_10': [0, 0, 0],  # 5-10倍赌注的次数
    'win_10_20': [0, 0, 0],  # 10-20倍赌注的次数
    'win_20_50': [0, 0, 0],  # 20-50倍赌注的次数
    'win_50_100': [0, 0, 0],  # 50-100倍赌注的次数
    'win_100_500': [0, 0, 0],  # 100-500倍赌注的次数
    'win_500_1000': [0, 0, 0],  # 500-1000倍赌注的次数
    'win_1000_5000': [0, 0, 0],  # 1000-5000倍赌注的次数
    'win_5000': [0, 0, 0],  # 5000倍以上赌注的次数
    
    'feature_0_10': 0,  # 特殊功能触发但赢额小于10倍赌注的次数
    'feature_0_10_win': 0,  # 特殊功能触发但赢额小于10倍赌注的总赢额
    'feature_1000_win': 0,  # 特殊功能触发且赢额大于等于500倍赌注的总赢额
    
    # 蟾蜍特色功能相关统计
    'toad_size_count': [0, 0, 0, 0, 0, 0],  # 蟾蜍大小统计(1-6)
    'toad_mul_count': [0 for i in range(50)],  # 蟾蜍乘数统计(0-49)
    'base_wild_buff_count': 0,  # 基础游戏中野生增益触发次数
    'win_buff_times': 0,  # 赢额增益触发次数
    'win_buff_win': 0,  # 赢额增益产生的总赢额
    'near_miss_count': 0,  # 接近触发但未触发特殊功能的次数("近似未中")
}


class TestCase(object):
    """测试用例类，包含测试方法和统计功能"""

    def test(self, test_time, total_bet, file_path, p_idx):
        """
        标准测试方法
        
        参数:
            test_time: 测试次数
            total_bet: 每次测试的总赌注
            file_path: 结果保存路径
            p_idx: 进程索引
        """
        limited_win = 25000  # 最大赢额限制(倍数)
        free_trigger_data = []  # 触发免费游戏的数据记录

        times = 0
        win_data = []  # 记录每次旋转的赢额倍数，用于计算方差
        while times < test_time:
            times += 1
            '''进度打印'''
            if times % (test_time / 20) == 0:
                if p_idx == 1:
                    print(f'Testing:\t{str(int(times / test_time * 100))}%')

            free_trigger = 0  # 免费游戏触发标记
            spin_total_win = 0  # 本次旋转总赢额
            
            # 执行基础游戏旋转
            base_result = Game_Slot.GameSlot().ng_spin(total_bet, Config.Base_Reel_Choose, Config.Base_Wild_Mul)

            # 累加基础游戏赢额
            spin_total_win += base_result['win_amount']
            base_wild_buff_pos = base_result['wild_buff_pos']  # 野生增益位置

            free_total_win = 0  # 免费游戏总赢额
            free_spin = base_result['free_spin']  # 获得的免费旋转次数
            
            # 如果触发免费游戏
            if free_spin > 0:
                # 复制蟾蜍信息并处理升级
                toad_info = copy.deepcopy(base_result['toad_info'])
                toad_info = Game_Slot.toad_move_upgrade(toad_info)  # 蟾蜍移动升级
                toad_info = Game_Slot.toad_move_upgrade(toad_info)  # 再次升级

                # 累加可能的额外免费旋转
                if toad_info['free_spin'] > 0:
                    free_spin += toad_info['free_spin']

                # 选择免费游戏轮盘集合
                choose_free_set_idx = Slot.rand_list(Config.Free_Set_Weight)
                free_set = Config.Free_Set[choose_free_set_idx]

            # 执行免费游戏
            use_fg_spin = 0  # 已使用的免费游戏次数
            low_win_spin = random.randint(1, 5)  # 随机选择低赢额旋转
            
            while free_spin > 0:
                free_spin -= 1
                use_fg_spin += 1

                # 确定是否应用赢额增益
                win_buff = 0
                if low_win_spin == use_fg_spin:
                    win_buff = 1
                    
                # 执行免费游戏旋转
                free_result = Game_Slot.GameSlot().fg_spin(total_bet, free_set, toad_info, free_spin, win_buff, free_total_win / total_bet)
                free_total_win += free_result['win_amount']  # 累加免费游戏赢额
                free_spin += free_result['free_spin']  # 累加额外获得的免费旋转
                toad_info = copy.deepcopy(free_result['toad_info'])  # 更新蟾蜍信息
                free_trigger += 1  # 增加免费游戏计数

                '''===免费游戏统计==='''
                data[Const.S_Free_Spin] += 1  # 增加免费旋转计数
                if free_result['win_amount'] > 0:
                    data[Const.S_Free_Win_Hit] += 1  # 增加免费游戏中奖计数

                # 统计最终蟾蜍状态
                toad_size = toad_info['toad_size']  # 蟾蜍大小
                toad_nul = toad_info['toad_mul']   # 蟾蜍乘数
                if free_spin == 0:  # 如果免费游戏结束
                    data['toad_size_count'][toad_size-1] += 1  # 记录最终蟾蜍大小
                    data['toad_mul_count'][toad_nul] += 1     # 记录最终蟾蜍乘数

            # 免费游戏特殊结果统计
            if free_trigger > 0:
                if free_total_win / total_bet >= 500:
                    data['feature_1000_win'] += free_total_win  # 高额赢金统计
                if free_total_win / total_bet < 10:
                    data['feature_0_10'] += 1  # 低额赢金次数
                    data['feature_0_10_win'] += free_total_win  # 低额赢金总额

            '''===基础游戏统计==='''
            data[Const.S_Bet] += total_bet  # 增加总赌注

            data[Const.S_Test_Time] += 1  # 增加测试次数

            # 野生增益统计
            if len(base_result['wild_buff_pos']) > 0:
                data['win_buff_times'] += 1  # 增加野生增益计数
                data['win_buff_win'] += base_result['win_amount']  # 累加野生增益赢额

            # 接近未中统计
            if base_result['near_miss'] > 0:
                data['near_miss_count'] += 1

            # 基础游戏中奖统计
            if base_result['win_amount'] > 0:
                data[Const.S_Base_Hit] += 1

            # 触发免费游戏统计
            if base_result['free_spin'] > 0:
                data[Const.S_Free_Hit] += 1

            # 累加基础游戏赢额
            data[Const.S_Base_Win] += base_result['win_amount']

            # 限制赢钱倍数 (防止极端大奖影响整体统计)
            if (base_result['win_amount'] + free_total_win) / total_bet > limited_win:
                data[Const.S_Win] += limited_win * total_bet  # 限制总赢额
                data[Const.S_Free_Win] += limited_win * total_bet - base_result['win_amount']  # 限制免费游戏赢额
                spin_total_win += limited_win * total_bet - base_result['win_amount']  # 限制本次旋转总赢额
                data['Max_Win_Times'] += 1  # 记录达到最大限制的次数
            else:
                data[Const.S_Win] += base_result['win_amount'] + free_total_win  # 累加总赢额
                data[Const.S_Free_Win] += free_total_win  # 累加免费游戏赢额
                spin_total_win += free_total_win  # 累加本次旋转总赢额

            # 进行赢额分布统计
            self.win_count(spin_total_win, total_bet, 0)  # 总体赢额统计
            if use_fg_spin > 0:
                self.win_count(spin_total_win, total_bet, 1)  # 免费游戏赢额统计

            # 野生增益赢额统计
            if len(base_wild_buff_pos) > 0:
                self.win_count(spin_total_win, total_bet, 2)  # 野生增益赢额统计
                data['base_wild_buff_count'] += 1  # 增加野生增益次数

            # 记录最大赢额
            if spin_total_win / total_bet > data['Max_Win'][p_idx]:
                data['Max_Win'][p_idx] = spin_total_win / total_bet

            # 记录赢额数据用于计算方差
            win_data.append(spin_total_win / total_bet)

        # 计算此进程的赢额方差
        data['Variance'][p_idx] = round(numpy.std(win_data),2)

        # 保存进程结果到文件
        file_name = file_path + '/' + str(p_idx) + '.txt'
        try:
            os.remove(file_name)
        except FileNotFoundError:
            pass
        with open(file_name, 'w', newline="") as f:
            f.write(json.dumps(data))


    def buy_test(self, test_time, total_bet, buy_type, file_path, p_idx):
        """
        购买特性测试方法
        
        参数:
            test_time: 测试次数
            total_bet: 每次测试的总赌注
            buy_type: 购买类型（1=普通免费游戏, 2=超级免费游戏）
            file_path: 结果保存路径
            p_idx: 进程索引
        """
        limited_win = 25000  # 最大赢额限制(倍数)
        # free_trigger_data = []

        times = 0
        win_data = []
        while times < test_time:
            times += 1
            '''进度打印'''
            if times % (test_time / 20) == 0:
                if p_idx == 1:
                    print(f'Testing:\t{str(int(times / test_time * 100))}%')

            free_trigger = 0
            spin_total_win = 0

            # 根据购买类型执行相应的旋转
            if buy_type == 1:
                base_result = Game_Slot.GameSlot().buy_ng_spin(total_bet, Reel_Choose=[], wild_set=Config.Buy_Base_Set, buy_free=1, buy_super=0)
            elif buy_type == 2:
                base_result = Game_Slot.GameSlot().buy_ng_spin(total_bet, Reel_Choose=[], wild_set=Config.Buy_Base_Set, buy_free=0, buy_super=1)
            else:
                print('spin false')
                base_result = Game_Slot.GameSlot().ng_spin(total_bet, Config.Base_Reel_Choose, Config.Base_Wild_Mul)


            spin_total_win += base_result['win_amount']



            # if p_idx == 1:
            #     print(base_result)

            free_total_win = 0
            free_spin = base_result['free_spin']
            if free_spin > 0:
                toad_info = copy.deepcopy(base_result['toad_info'])
                toad_info = Game_Slot.toad_move_upgrade(toad_info)

                if toad_info['free_spin'] > 0:
                    free_spin += toad_info['free_spin']
                # if len(free_trigger_data) < 20 and p_idx == 1:
                #     free_trigger_data.append(base_result)
                # if p_idx == 1:
                #     print('\n\n\n')
                if buy_type == 1:
                    choose_free_set_idx = Slot.rand_list(Config.Buy_Free_Set_Weight)
                    free_set = Config.Buy_Set[choose_free_set_idx-1]
                elif buy_type == 2:
                    choose_free_set_idx = Slot.rand_list(Config.Buy_Super_Set_Weight)
                    free_set = Config.Buy_Super_Set[choose_free_set_idx-1]
                else:
                    print('free set choose false')
                    free_set = {}

            use_fg_spin = 0
            low_win_spin = random.randint(1, 5)
            while free_spin > 0:
                free_spin -= 1
                use_fg_spin += 1

                win_buff = 0
                if low_win_spin == use_fg_spin:
                    win_buff = 1

                if buy_type == 1:
                    free_result = Game_Slot.GameSlot().fg_spin(total_bet, free_set, toad_info, free_spin, win_buff, free_total_win / total_bet, buy_free=1,buy_super=0)
                elif buy_type == 2:
                    free_result = Game_Slot.GameSlot().fg_spin(total_bet, free_set, toad_info, free_spin, win_buff, free_total_win / total_bet, buy_free=0, buy_super=1)

                free_total_win += free_result['win_amount']
                free_spin += free_result['free_spin']
                toad_info = copy.deepcopy(free_result['toad_info'])
                free_trigger += 1

                # if p_idx == 1:
                #     print(free_result)
                #     print(use_fg_spin)
                #     print(free_result['reel_info']['item_list'])
                #     for i in free_result['win_info']['lines']:
                #         print(i)
                #     print(free_result['toad_info'])
                #     print(free_result['win_amount'])
                #     print('\n')

                '''===free 统计==='''
                data[Const.S_Free_Spin] += 1
                if free_result['win_amount'] > 0:
                    data[Const.S_Free_Win_Hit] += 1

                toad_size = toad_info['toad_size']
                toad_nul = toad_info['toad_mul']
                if free_spin == 0:
                    data['toad_size_count'][toad_size-1] += 1
                    data['toad_mul_count'][toad_nul] += 1


            if free_trigger > 0:
                if free_total_win / total_bet >= 1000:
                    data['feature_1000_win'] += free_total_win
                if free_total_win / total_bet < 10:
                    data['feature_0_10'] += 1
                    data['feature_0_10_win'] += free_total_win

            '''===base统计==='''
            if buy_type == 1:
                data[Const.S_Bet] += total_bet * 100
            elif buy_type == 2:
                data[Const.S_Bet] += total_bet * 500


            data[Const.S_Test_Time] += 1


            if base_result['win_amount'] > 0:
                data[Const.S_Base_Hit] += 1

            if base_result['free_spin'] > 0:
                data[Const.S_Free_Hit] += 1


            data[Const.S_Base_Win] += base_result['win_amount']

            # 限制赢钱倍数
            if (base_result['win_amount'] + free_total_win) / total_bet > limited_win:
                data[Const.S_Win] += limited_win * total_bet
                data[Const.S_Free_Win] += limited_win * total_bet - base_result['win_amount']
                spin_total_win += limited_win * total_bet - base_result['win_amount']
                data['Max_Win_Times'] += 1

            else:
                data[Const.S_Win] += base_result['win_amount'] + free_total_win
                data[Const.S_Free_Win] += free_total_win
                spin_total_win += free_total_win



            # # 累加单次spin + free 的总赢金
            # data[Const.S_Win] += base_result['win_amount'] + free_total_win
            # data[Const.S_Free_Win] += free_total_win
            # spin_total_win += free_total_win

            self.win_count(spin_total_win, total_bet, 0)
            if free_total_win > 0:
                self.win_count(spin_total_win, total_bet, 1)
            if spin_total_win / total_bet > data['Max_Win'][p_idx]:
                data['Max_Win'][p_idx] = spin_total_win / total_bet

            win_data.append(spin_total_win / total_bet)

        data['Variance'][p_idx] = round(numpy.std(win_data),2)
        file_name = file_path + '/' + str(p_idx) + '.txt'

        # if p_idx == 1:
        #     print(trigger_data)

        try:
            os.remove(file_name)
        except FileNotFoundError:
            pass
        with open(file_name, 'w', newline="") as f:
            f.write(json.dumps(data))


    def win_count(self, spin_total_win, total_bet, idx):
        """
        赢额分布统计方法
        
        参数:
            spin_total_win: 本次旋转总赢额
            total_bet: 每次测试的总赌注
            idx: 统计类别索引 (0=总体, 1=免费游戏, 2=野生增益)
        """
        if spin_total_win / total_bet == 0:
            data['win_0'][idx] += 1
        elif 0 < spin_total_win / total_bet <= 1:
            data['win_0_1'][idx] += 1
        elif 1 < spin_total_win / total_bet <= 2:
            data['win_1_2'][idx] += 1
        elif 2 < spin_total_win / total_bet <= 3:
            data['win_2_3'][idx] += 1
        elif 3 < spin_total_win / total_bet <= 4:
            data['win_3_4'][idx] += 1
        elif 4 < spin_total_win / total_bet <= 5:
            data['win_4_5'][idx] += 1
        elif 5 < spin_total_win / total_bet <= 10:
            data['win_5_10'][idx] += 1
        elif 10 < spin_total_win / total_bet <= 20:
            data['win_10_20'][idx] += 1
        elif 20 < spin_total_win / total_bet <= 50:
            data['win_20_50'][idx] += 1
        elif 50 < spin_total_win / total_bet <= 100:
            data['win_50_100'][idx] += 1
        elif 100 < spin_total_win / total_bet <= 500:
            data['win_100_500'][idx] += 1
        elif 500 < spin_total_win / total_bet <= 1000:
            data['win_500_1000'][idx] += 1
        elif 1000 < spin_total_win / total_bet <= 5000:
            data['win_1000_5000'][idx] += 1
        elif 5000 < spin_total_win / total_bet:
            data['win_5000'][idx] += 1



def run(test_time, total_bet, file_path, p_idx):
    """
    运行测试函数，每个进程调用此函数
    
    参数:
        test_time: 测试次数
        total_bet: 总赌注
        file_path: 结果保存路径
        p_idx: 进程索引
    """
    # buy_type = 1  # 设置购买类型: 2表示购买超级免费游戏
    TestCase().test(test_time, total_bet, file_path, p_idx)  # 标准测试(已注释)
    # TestCase().buy_test(test_time, total_bet, buy_type, file_path, p_idx)  # 购买特性测试
    print(str(p_idx) + ": Over")  # 打印进程完成信息


if __name__ == '__main__':
    # 程序入口：配置测试参数并启动多进程测试
    start_time = datetime.datetime.now()
    test_time = 1000000
    total_bet = 100

    file_list = []
    unit_times = test_time / 10


    data_store_dir = '/home/wqh/project/Slot013Numeric/Slot013Python/data/' + start_time.strftime("%Y_%m_%d_%H_%M_%S")
    os.mkdir(data_store_dir)

    # 创建进程列表（1-25 编号）
    processes = [
        multiprocessing.Process(
            target=run,
            args=(unit_times, total_bet, data_store_dir, i)  # i 为 1-25 的编号
        )
        for i in range(1, 11)  # 生成 25 个进程
    ]

    # 启动所有进程
    for p in processes:
        p.start()

    # 等待所有进程完成
    for p in processes:
        p.join()

    file_list = os.listdir(data_store_dir)

    data_list = []
    for file in file_list:
        file_path = os.path.join(data_store_dir, file)
        with open(file_path, 'r') as f:
            contents = f.read()
            data_list.append(json.loads(contents))
    sum_data = copy.deepcopy(data)
    for k, v in sum_data.items():
        for sub_data in data_list:
            # int 相加
            if isinstance(sub_data[k], int) or isinstance(sub_data[k], float):
                sum_data[k] += sub_data[k]
            # list 相加
            elif isinstance(sub_data[k], list):
                for x in range(len(sum_data[k])):
                    sum_data[k][x] += sub_data[k][x]
            # dict 相加
            elif isinstance(sub_data[k], dict):
                for k2 in sum_data[k].keys():
                    if isinstance(sum_data[k][k2], list):
                        for i in range(len(sum_data[k][k2])):
                            sum_data[k][k2][i] += sub_data[k][str(k2)][i]
                    elif isinstance(sum_data[k][k2], int) or isinstance(sum_data[k][k2], float):
                        sum_data[k][k2] += sub_data[k][str(k2)]

    print(f"Total RTP：{sum_data[Const.S_Win] / sum_data[Const.S_Bet]}")
    print('=============================')
    print(f'Base RTP：{sum_data[Const.S_Base_Win] / sum_data[Const.S_Bet]}')
    print(f'Base Hit Rate：{sum_data[Const.S_Base_Hit] / sum_data[Const.S_Test_Time]}')
    try:
        print(f'Near Miss 间隔: ', sum_data[Const.S_Test_Time] / sum_data['near_miss_count'])
    except ZeroDivisionError:
        pass

    try:
        print(f'天降横财 RTP:', sum_data['win_buff_win'] / sum_data[Const.S_Bet])
        print(f'天降横财的平均倍数:', sum_data['win_buff_win'] / sum_data['win_buff_times'] / total_bet)
        print(f'天降横财的间隔:', sum_data[Const.S_Test_Time] / sum_data['win_buff_times'])
        print('Base 天降概率:', sum_data['base_wild_buff_count'] / sum_data[Const.S_Test_Time])
    except:
        pass

    print('=============================')
    try:
        print(f'Free RTP：{(sum_data[Const.S_Free_Win]) / sum_data[Const.S_Bet]}')
        print(f'Free触发间隔：{sum_data[Const.S_Test_Time] / sum_data[Const.S_Free_Hit]}')

        print(f'Free平均倍数：{(sum_data[Const.S_Free_Win]) / sum_data[Const.S_Free_Hit] / total_bet}')
        print(f'Free hit rate：{sum_data[Const.S_Free_Win_Hit] / sum_data[Const.S_Free_Spin]}')
        print(f'Free平均次数：{sum_data[Const.S_Free_Spin] / sum_data[Const.S_Free_Hit]}')
        print('Free win < 10 total bet: 占比/ RTP', sum_data['feature_0_10'] / sum_data[Const.S_Free_Hit], sum_data['feature_0_10_win'] / sum_data[Const.S_Bet])
        print('Free win >= 500 total bet: RTP', sum_data['feature_1000_win'] / sum_data[Const.S_Bet])
    except:
        pass

    print('=============================')
    max_win_list = []
    for i in sum_data['Max_Win']:
        if i != 0:
            max_win_list.append(i)
    print('Max_Win:', max_win_list)
    print('Max_Win Time / Test Times:', sum_data['Max_Win_Times'], '/', sum_data[Const.S_Test_Time])

    sample_variance = []
    for i in sum_data['Variance']:
        if i > 0:
            sample_variance.append(i)

    total_variance = round(sum(sample_variance) / len(sample_variance), 4)

    print('方差 Variance:',total_variance)
    win_table = pt.PrettyTable()
    win_table.field_names = ["Win Mul", 'total', 'free', 'base天降']
    for k in ['win_0', 'win_0_1', 'win_1_2', 'win_2_3', 'win_3_4', 'win_4_5', 'win_5_10', 'win_10_20', 'win_20_50', 'win_50_100', 'win_100_500', 'win_500_1000', 'win_1000_5000', 'win_5000']:
        try:
            win_table.add_row([k, round(sum_data[k][0] / sum_data[Const.S_Test_Time], 4), round(sum_data[k][1] / sum_data[Const.S_Free_Hit], 4), round(sum_data[k][2] / sum_data['base_wild_buff_count'], 4)])
        except ZeroDivisionError:
            win_table.add_row([k, round(sum_data[k][0] / sum_data[Const.S_Test_Time], 4), round(sum_data[k][1] / sum_data[Const.S_Free_Hit], 4), 0])

    print(win_table)

    print('=============================')
    toad_size_table = pt.PrettyTable()
    toad_size_table.field_names = ["toad size", 'percent']
    for k in range(0, 6):
        toad_size_table.add_row([k+1, round(sum_data['toad_size_count'][k] / sum(sum_data['toad_size_count']), 5)])

    print(toad_size_table)
    print('=============================')
    toad_mul_table = pt.PrettyTable()
    toad_mul_table.field_names = ["toad Mul", 'percent']
    for k in range(len(sum_data['toad_mul_count'])):
        if sum_data['toad_mul_count'][k] != 0:
            toad_mul_table.add_row([k, round(sum_data['toad_mul_count'][k] / sum(sum_data['toad_mul_count']), 5)])

    print(toad_mul_table)
    end_time = datetime.datetime.now()
    spend_time = (end_time - start_time).seconds
    print(f"Test_Time：{sum_data[Const.S_Test_Time]}")
    print('Spend Time：' + str(spend_time) + 's')
