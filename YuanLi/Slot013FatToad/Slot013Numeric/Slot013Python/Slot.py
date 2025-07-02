import random
import copy
import Const as Const


class DealReel(object):
    """
    处理reel，转换原始卷轴数据为可使用的格式
    卷轴数据通常包含符号及其权重，需要转换为便于随机抽取的累积权重形式
    """
    def reel_strip(self, reel_set):
        """
        处理整个卷轴集合，将每个卷轴的原始数据转换为处理后的格式
        
        参数:
            reel_set: 卷轴集合，包含多个卷轴的数据
            
        返回:
            处理后的卷轴集合
        """
        for reel in reel_set:
            deal_reel = []
            strip = reel[Const.C_ReelStrip]
            for single_reel in strip:
                deal_reel.append(self.deal_singal_reel(single_reel))
            reel[Const.R_Dealed_ReelStrip] = deal_reel
        return reel_set

    def deal_singal_reel(self, single_reel):
        """
        处理单个卷轴的数据，将原始的[符号,权重]对转换为累积权重形式
        
        参数:
            single_reel: 单个卷轴的原始数据
            
        返回:
            包含符号列表和累积权重列表的字典
        """
        symbol_list = []
        weight_list = []
        accumulate_weight = 0
        for i in single_reel:
            accumulate_weight += i[1]
            symbol_list.append(i[0])
            weight_list.append(accumulate_weight)
        return {
            Const.R_Symbol_List: symbol_list,
            Const.R_Weight_List: weight_list
        }


class GetReel(object):
    """
    获取卷轴的随机结果
    负责从卷轴数据中随机生成游戏中显示的符号组合
    """

    def __init__(self, reel_set, shape):
        """
        初始化卷轴生成器
        
        参数:
            reel_set: 卷轴集合数据
            shape: 各列显示的符号数量，如[3,3,3,3,3]表示5列各显示3个符号
        """
        self.reel_set = reel_set
        self.reel_strips = reel_set[Const.C_ReelStrip]
        self.shape = shape


    def get_pos(self):
        """
        为每列卷轴随机生成一个起始位置
        
        返回:
            包含各列起始位置的列表
        """
        pos = []
        for single_reel in self.reel_strips:
            # 随机选择一个位置作为该列卷轴的起始点
            sym_num = len(single_reel)
            ra = random.randint(0, sym_num-1)
            pos.append(ra)
        return pos


    def get_reel(self, test_pos=[]):
        """
        获取实际的卷轴结果
        
        参数:
            test_pos: 可选的测试位置，用于调试或固定结果
            
        返回:
            包含卷轴结果、位置和掉落数量的字典
        """
        if len(test_pos) > 0:
            pos = test_pos
        else:
            pos = self.get_pos()

        reel = []

        # 根据起始位置和形状生成各列的符号
        for reel_idx in range(len(pos)):
            # 将卷轴复制一次以处理循环边界问题
            long_reel = self.reel_strips[reel_idx] + self.reel_strips[reel_idx]
            reel.append(long_reel[pos[reel_idx]:pos[reel_idx] + self.shape[reel_idx]])

        # 处理神秘符号(Mystery Symbol)逻辑
        if self.reel_set[Const.C_Mystery] is not False:
            # 类型1: 每列的神秘符号转变为同一个符号
            if self.reel_set[Const.C_Mystery_Type] == 1:
                for x in range(len(reel)):
                    mystery_change = rand_list(self.reel_set[Const.C_Mystery_Weight])
                    for y in range(len(reel[x])):
                        if reel[x][y] == self.reel_set[Const.C_Mystery]:
                            reel[x][y] = mystery_change
            # 类型2: 所有神秘符号转变为同一个符号
            elif self.reel_set[Const.C_Mystery_Type] == 2:
                mystery_change = rand_list(self.reel_set[Const.C_Mystery_Weight])
                for x in range(len(reel)):
                    for y in range(len(reel[x])):
                        if reel[x][y] == self.reel_set[Const.C_Mystery]:
                            reel[x][y] = mystery_change
            else:
                print('Mystery type error')

        # 如果需要，对卷轴进行随机排列
        if self.reel_set[Const.C_Shuffle] is True:
            random.shuffle(reel)
        return {
            'item_list': reel,            # 卷轴上的符号列表
            Const.R_Reel_Pos: pos,        # 各列的起始位置
            Const.R_Reel_Drop_Num: copy.deepcopy(self.shape)  # 各列的符号数量
        }


    def get_drop_reel(self, reel_info):
        """
        获取掉落式卷轴的新符号
        用于处理消除后的新符号掉落逻辑
        
        参数:
            reel_info: 当前卷轴信息
            
        返回:
            包含新掉落符号的卷轴信息
        """
        pos = reel_info[Const.R_Reel_Pos]
        drop_reel = []

        # 计算新的掉落位置
        drop_pos = copy.deepcopy(pos)
        for x in range(len(pos)):
            drop_pos[x] += reel_info[Const.R_Reel_Drop_Num][x]

        # 获取每列的新符号
        for reel_idx in range(len(drop_pos)):
            long_reel = self.reel_strips[reel_idx] + self.reel_strips[reel_idx]
            drop_reel.append(long_reel[drop_pos[reel_idx]:drop_pos[reel_idx] + self.shape[reel_idx]])

        # 更新掉落数量
        for x in range(len(pos)):
            reel_info[Const.R_Reel_Drop_Num][x] += self.shape[x]
        return {
            'drop_item_list': drop_reel,  # 新掉落的符号列表
            Const.R_Reel_Pos: pos,        # 各列的起始位置
            Const.R_Reel_Drop_Num: copy.deepcopy(reel_info[Const.R_Reel_Drop_Num])  # 更新后的掉落数量
        }


class StandardLineEvaluator(object):
    """
    处理reel结果，计算线上赢钱
    标准老虎机的赢钱线评估器，计算各条赢钱线上的获胜组合和奖金
    """
    def __init__(self, totalbet, spinreel, payline, paytable, betLine, wild_sub, linesym, wilds, wild):
        """
        初始化评估器
        
        参数:
            totalbet: 总下注额
            spinreel: 当前旋转结果，即卷轴上显示的符号
            payline: 赢钱线定义，每条线指定各列的行位置
            paytable: 赢钱表，定义各符号组合的赔率
            betLine: 下注线数量
            wild_sub: 可被百搭符号替代的普通符号列表
            linesym: 可形成连线的符号列表
            wilds: 百搭符号列表
            wild: 主百搭符号
        """
        self.totalbet = totalbet  # 押注
        self.spinreel = spinreel  # 盘面结果

        self.paytable = paytable  # 赢钱表
        self.payline = payline    # 赢钱线定义

        self.wildsub = wild_sub   # wild可替代图标数组
        self.line_sym = linesym   # 连线图标
        self.wilds = wilds        # wild类图标
        self.wild = wild          # 主wild符号
        self.betLine = betLine    # 押注线数

        self.linebet = totalbet / betLine  # 每条线基础押注额

    def GetAllLine(self):
        """
        获取所有赢钱线上的符号组合
        
        返回:
            包含各线信息的列表，每条线包含ID、符号组合和位置信息
        """
        line_id = 0
        line_result = []
        for line in self.payline:
            col = 0
            line_combo = []  # 该线上的符号组合
            line_pos = []    # 该线上的符号位置
            
            # 根据赢钱线定义收集每列的符号
            for row in line:
                line_combo.append(self.spinreel[col][row])
                line_pos.append([col, row])
                col += 1
            line_result.append({
                Const.R_Line_Id: line_id, 
                Const.R_Line_Combo: line_combo, 
                Const.R_Line_Pos: line_pos
            })
            line_id += 1
        return line_result

    def evaluateLine(self, oneline):
        """
        评估单条线的获胜情况
        
        参数:
            oneline: 单条线的信息
            
        返回:
            添加了获胜信息的线条数据
        """
        line_combo = oneline[Const.R_Line_Combo]
        sym_long = 0      # 符号连续数量
        wild_long = 0     # wild连续数量
        kind = line_combo[0]  # 连线符号类型
        line_mul = 0      # 赢钱倍数
        long = 0          # 连线长度

        # 只有当首个符号是可连线符号时才进行评估
        if line_combo[0] in self.line_sym:
            # 情况1: 首个符号是普通符号(可被wild替代)
            if line_combo[0] in self.wildsub:
                # 计算从左到右连续的相同符号数量(包括wild替代)
                for x in range(len(line_combo)):
                    if line_combo[x] == line_combo[0] or line_combo[x] in self.wilds:
                        sym_long += 1
                    else:
                        break

                kind = line_combo[0]
                line_mul = self.paytable[kind][sym_long - 1]
                long = sym_long

            # 情况2: 首个符号是wild
            elif line_combo[0] in self.wilds:
                # 计算从左到右连续的wild符号数量
                for x in range(len(line_combo)):
                    if line_combo[x] in self.wilds:
                        wild_long += 1
                    else:
                        break

                # 如果wild不是一直连到最后
                if wild_long < len(line_combo):
                    # 获取wild之后的第一个普通符号
                    normal_sym = line_combo[wild_long]

                    # 如果这个普通符号可以连线
                    if normal_sym in self.line_sym:
                        # 计算该符号的连线长度(包括前面的wild)
                        for x in range(len(line_combo)):
                            if line_combo[x] == line_combo[wild_long] or line_combo[x] in self.wilds:
                                sym_long += 1
                            else:
                                break

                        # 比较普通符号连线和纯wild连线哪个赔率更高
                        if self.paytable[normal_sym][sym_long - 1] >= self.paytable[self.wild][wild_long - 1]:
                            line_mul = self.paytable[normal_sym][sym_long - 1]
                            long = sym_long
                            kind = normal_sym
                        else:
                            line_mul = self.paytable[self.wild][wild_long - 1]
                            long = wild_long
                            kind = self.wild
                    else:
                        # 如果wild后的符号不是连线符号，则只计算wild连线
                        line_mul = self.paytable[self.wild][wild_long - 1]
                        long = wild_long
                        kind = self.wild
                else:
                    # 全部都是wild的情况
                    line_mul = self.paytable[self.wild][wild_long - 1]
                    long = wild_long
                    kind = self.wild

        # 更新线条信息
        oneline[Const.R_Line_Kind] = kind              # 连线的符号类型
        oneline[Const.R_Line_Mul] = line_mul           # 赢钱倍数
        oneline[Const.R_Line_Win] = line_mul * self.linebet  # 赢钱金额
        oneline[Const.R_Line_Pos] = oneline[Const.R_Line_Pos][:long]  # 连线位置
        oneline[Const.R_Line_Long] = long              # 连线长度

        return oneline

    def evaluate(self):
        """
        评估所有线的获胜情况(基本游戏)
        
        返回:
            包含所有获胜线信息的结果
        """
        line_result = {}
        line_ifo = self.GetAllLine()
        line_result[Const.R_Line] = []
        win_pos_list = []  # 所有获胜位置
        win_amount = 0     # 总赢钱金额
        wild_2_active = 0  # 特殊wild激活状态
        
        for simple_line in line_ifo:
            oneline = self.evaluateLine(simple_line)

            # 如果这条线赢钱了
            if oneline[Const.R_Line_Win] > 0:
                long = oneline[Const.R_Line_Long]

                # 处理wild乘数逻辑 (特殊wild符号12,13,15会提供额外乘数)
                wild_mul = 0
                for i in range(long):
                    if oneline[Const.R_Line_Combo][i] in [12, 13, 15]:
                        wild_mul += oneline[Const.R_Line_Combo][i] - 10
                        wild_2_active = 1

                # 应用wild乘数
                if wild_mul != 0:
                    oneline[Const.R_Line_Win] = oneline[Const.R_Line_Win] * wild_mul

                win_amount += oneline[Const.R_Line_Win]
                line_result[Const.R_Line].append(oneline)

                # 收集所有获胜位置
                for pos in oneline[Const.R_Line_Pos]:
                    if pos not in win_pos_list:
                        win_pos_list.append(pos)

        # 汇总结果
        line_result[Const.R_Line_WinAmount] = win_amount
        line_result[Const.R_Win_Amount] = win_amount
        line_result[Const.R_Win_Pos_List] = win_pos_list
        line_result['wild_2_active'] = wild_2_active
        return line_result

    def free_evaluate(self, toad_wild):
        """
        评估免费游戏中所有线的获胜情况
        免费游戏中有额外的蟾蜍wild乘数逻辑
        
        参数:
            toad_wild: 蟾蜍wild符号ID
            
        返回:
            包含所有获胜线信息的结果
        """
        line_result = {}
        line_ifo = self.GetAllLine()
        line_result[Const.R_Line] = []
        win_pos_list = []
        win_amount = 0
        wild_2_active = False
        
        for simple_line in line_ifo:
            oneline = self.evaluateLine(simple_line)

            if oneline[Const.R_Line_Win] > 0:
                long = oneline[Const.R_Line_Long]

                # 处理wild乘数逻辑 (免费游戏中12-19的wild提供额外乘数)
                wild_mul = 0
                for i in range(long):
                    if 12 <= oneline[Const.R_Line_Combo][i] < 20:
                        wild_mul += oneline[Const.R_Line_Combo][i] - 10
                        wild_2_active = True

                # 处理蟾蜍wild额外乘数
                if toad_wild in oneline[Const.R_Line_Combo][:long]:
                    wild_mul += toad_wild - 20

                # 应用wild乘数
                if wild_mul != 0:
                    oneline[Const.R_Line_Win] = oneline[Const.R_Line_Win] * wild_mul

                win_amount += oneline[Const.R_Line_Win]
                line_result[Const.R_Line].append(oneline)

                # 收集所有获胜位置
                for pos in oneline[Const.R_Line_Pos]:
                    if pos not in win_pos_list:
                        win_pos_list.append(pos)

        # 汇总结果
        line_result[Const.R_Line_WinAmount] = win_amount
        line_result[Const.R_Win_Amount] = win_amount
        line_result[Const.R_Win_Pos_List] = win_pos_list
        line_result['wild_2_active'] = wild_2_active
        return line_result


def rand_list(one_list):
    """
    从带权重的列表中随机选择一个元素
    
    参数:
        one_list: 形如[[元素1,权重1],[元素2,权重2],...]的列表
        
    返回:
        随机选中的元素
    """
    total_weight = 0
    for i in one_list:
        total_weight += i[1]
    ra = random.randint(0, total_weight - 1)
    curr_sum = 0

    kind = None
    for k in one_list:
        curr_sum = curr_sum + k[1]
        if ra < curr_sum:
            kind = k[0]
            break
    return kind


def rand_dict(one_dict):
    """
    从带权重的字典中随机选择一个键
    
    参数:
        one_dict: 形如{键1:权重1, 键2:权重2, ...}的字典
        
    返回:
        随机选中的键
    """
    total_weight = sum(one_dict.values())
    ra = random.randint(0, total_weight - 1)

    curr_sum = 0
    keys = one_dict.keys()
    kind = None
    for k in keys:
        curr_sum = curr_sum + one_dict[k]
        if ra < curr_sum:
            kind = k
            break
    return kind
