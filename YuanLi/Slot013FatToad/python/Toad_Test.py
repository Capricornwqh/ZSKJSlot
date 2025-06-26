import multiprocessing
import datetime
import json
import copy
import os
import random

import numpy

import Toad_Slot as Game_Slot
import Toad_Config as Config
import Slot
import Const as Const
import base_free_trigger_data
import prettytable as pt


data = {
    Const.S_Test_Time: 0,

    Const.S_Bet: 0,
    Const.S_Win: 0,

    Const.S_Base_Win: 0,
    Const.S_SC_Win: 0,

    Const.S_Base_Hit: 0,
    Const.S_Base_Sym_Win: {
        Config.H1: [0, 0, 0, 0, 0, 0],
        Config.H2: [0, 0, 0, 0, 0, 0],
        Config.H3: [0, 0, 0, 0, 0, 0],
        Config.H4: [0, 0, 0, 0, 0, 0],
        Config.LA: [0, 0, 0, 0, 0, 0],
        Config.LK: [0, 0, 0, 0, 0, 0],
        Config.LQ: [0, 0, 0, 0, 0, 0],
        Config.LJ: [0, 0, 0, 0, 0, 0],
    },

    Const.S_Free_Hit: 0,
    Const.S_Free_Win_Hit: 0,
    Const.S_Free_Win: 0,
    Const.S_Free_Spin: 0,


    Const.S_Feature_Hit: 0,
    'Max_Win': [0 for _ in range(40)],
    'Max_Win_Times': 0,
    'Variance': [0 for _ in range(40)],
    'win_0': [0, 0, 0],
    'win_0_1': [0, 0, 0],
    'win_1_2': [0, 0, 0],
    'win_2_3': [0, 0, 0],
    'win_3_4': [0, 0, 0],
    'win_4_5': [0, 0, 0],
    'win_5_10': [0, 0, 0],
    'win_10_20': [0, 0, 0],
    'win_20_50': [0, 0, 0],
    'win_50_100': [0, 0, 0],
    'win_100_500': [0, 0, 0],
    'win_500_1000': [0, 0, 0],
    'win_1000_5000': [0, 0, 0],
    'win_5000': [0, 0, 0],
    'feature_0_10': 0,
    'feature_0_10_win': 0,
    'feature_1000_win': 0,
    'toad_size_count': [0, 0, 0, 0, 0, 0],
    'toad_mul_count': [0 for i in range(50)],
    'base_wild_buff_count': 0,
    'win_buff_times': 0,
    'win_buff_win': 0,
    'near_miss_count': 0,
}


class TestCase(object):


    def test(self, test_time, total_bet, file_path, p_idx):
        limited_win = 25000
        free_trigger_data = []

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
            base_result = Game_Slot.GameSlot().ng_spin(total_bet, Config.Base_Reel_Choose, Config.Base_Wild_Mul)

            # base_result = random.choice(base_free_trigger_data.free_trigger)

            spin_total_win += base_result['win_amount']
            base_wild_buff_pos = base_result['wild_buff_pos']


            # if p_idx == 1:
            #     print(base_result)

            free_total_win = 0
            free_spin = base_result['free_spin']
            if free_spin > 0:
                toad_info = copy.deepcopy(base_result['toad_info'])
                toad_info = Game_Slot.toad_move_upgrade(toad_info)

                toad_info = Game_Slot.toad_move_upgrade(toad_info)

                if toad_info['free_spin'] > 0:
                    free_spin += toad_info['free_spin']

                # if len(free_trigger_data) < 1000 and p_idx == 1:
                #     free_trigger_data.append(base_result)
                # else:
                #     times = test_time
                # if p_idx == 1:
                #     print('\n\n\n')

                choose_free_set_idx = Slot.rand_list(Config.Free_Set_Weight)
                free_set = Config.Free_Set[choose_free_set_idx]

            use_fg_spin = 0
            low_win_spin = random.randint(1, 5)
            while free_spin > 0:
                free_spin -= 1
                use_fg_spin += 1

                win_buff = 0
                if low_win_spin == use_fg_spin:
                    win_buff = 1
                free_result = Game_Slot.GameSlot().fg_spin(total_bet, free_set, toad_info, free_spin, win_buff, free_total_win / total_bet)
                free_total_win += free_result['win_amount']
                free_spin += free_result['free_spin']
                toad_info = copy.deepcopy(free_result['toad_info'])
                free_trigger += 1

                # if p_idx == 1:
                #     print(use_fg_spin)
                #     print(free_result['reel_info']['item_list'])
                #     # for i in free_result['win_info']['lines']:
                #     #     print(i)
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
                if free_total_win / total_bet >= 500:
                    data['feature_1000_win'] += free_total_win
                if free_total_win / total_bet < 10:
                    data['feature_0_10'] += 1
                    data['feature_0_10_win'] += free_total_win

            '''===base统计==='''
            data[Const.S_Bet] += total_bet

            data[Const.S_Test_Time] += 1

            if len(base_result['wild_buff_pos']) > 0:
                data['win_buff_times'] += 1
                data['win_buff_win'] += base_result['win_amount']

            if base_result['near_miss'] > 0:
                data['near_miss_count'] += 1

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
            if use_fg_spin > 0:
                self.win_count(spin_total_win, total_bet, 1)

            if len(base_wild_buff_pos) > 0:
                self.win_count(spin_total_win, total_bet, 2)
                data['base_wild_buff_count'] += 1

            if spin_total_win / total_bet > data['Max_Win'][p_idx]:
                data['Max_Win'][p_idx] = spin_total_win / total_bet

            win_data.append(spin_total_win / total_bet)

        data['Variance'][p_idx] = round(numpy.std(win_data),2)

        file_name = file_path + '/' + str(p_idx) + '.txt'

        # if p_idx == 1:
        #     for i in free_trigger_data:
        #         print(i, ',')

        try:
            os.remove(file_name)
        except FileNotFoundError:
            pass
        with open(file_name, 'w', newline="") as f:
            f.write(json.dumps(data))


    def buy_test(self, test_time, total_bet, buy_type, file_path, p_idx):
        limited_win = 25000
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
    buy_type = 2
    # TestCase().test(test_time, total_bet, file_path, p_idx)
    TestCase().buy_test(test_time, total_bet, buy_type, file_path, p_idx)
    print(str(p_idx) + ": Over")


if __name__ == '__main__':

    start_time = datetime.datetime.now()
    test_time = 10000000
    total_bet = 100

    file_list = []
    unit_times = test_time / 25


    data_store_dir = 'E:\\Run_Data\\' + start_time.strftime("%Y_%m_%d_%H_%M_%S")
    os.mkdir(data_store_dir)

    # 创建进程列表（1-25 编号）
    processes = [
        multiprocessing.Process(
            target=run,
            args=(unit_times, total_bet, data_store_dir, i)  # i 为 1-25 的编号
        )
        for i in range(1, 26)  # 生成 25 个进程
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
