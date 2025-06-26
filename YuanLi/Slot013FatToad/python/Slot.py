import random
import copy
import Const as Const


class DealReel(object):
    """
    处理reel，转换
    """
    def reel_strip(self, reel_set):
        for reel in reel_set:
            deal_reel = []
            strip = reel[Const.C_ReelStrip]
            for single_reel in strip:
                deal_reel.append(self.deal_singal_reel(single_reel))
            reel[Const.R_Dealed_ReelStrip] = deal_reel
        return reel_set

    def deal_singal_reel(self, single_reel):
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
    get 卷轴随机结果
    """

    def __init__(self, reel_set, shape):
        self.reel_set = reel_set

        self.reel_strips = reel_set[Const.C_ReelStrip]
        self.shape = shape


    def get_pos(self):
        pos = []
        for single_reel in self.reel_strips:

            sym_num = len(single_reel)
            ra = random.randint(0, sym_num-1)
            pos.append(ra)
        return pos



    def get_reel(self, test_pos=[]):
        if len(test_pos) > 0:
            pos = test_pos
        else:
            pos = self.get_pos()

        reel = []

        for reel_idx in range(len(pos)):
            long_reel = self.reel_strips[reel_idx] + self.reel_strips[reel_idx]
            reel.append(long_reel[pos[reel_idx]:pos[reel_idx] + self.shape[reel_idx]])

        if self.reel_set[Const.C_Mystery] is not False:

            if self.reel_set[Const.C_Mystery_Type] == 1:
                for x in range(len(reel)):
                    mystery_change = rand_list(self.reel_set[Const.C_Mystery_Weight])
                    for y in range(len(reel[x])):
                        if reel[x][y] == self.reel_set[Const.C_Mystery]:
                            reel[x][y] = mystery_change
            elif self.reel_set[Const.C_Mystery_Type] == 2:
                mystery_change = rand_list(self.reel_set[Const.C_Mystery_Weight])
                for x in range(len(reel)):
                    for y in range(len(reel[x])):
                        if reel[x][y] == self.reel_set[Const.C_Mystery]:
                            reel[x][y] = mystery_change
            else:
                print('Mystery type error')


        if self.reel_set[Const.C_Shuffle] is True:
            random.shuffle(reel)
        return {
            'item_list': reel,
            Const.R_Reel_Pos: pos,
            Const.R_Reel_Drop_Num: copy.deepcopy(self.shape)
        }



    def get_drop_reel(self, reel_info):
        pos = reel_info[Const.R_Reel_Pos]
        drop_reel = []

        drop_pos = copy.deepcopy(pos)
        for x in range(len(pos)):
            drop_pos[x] += reel_info[Const.R_Reel_Drop_Num][x]

        for reel_idx in range(len(drop_pos)):
            long_reel = self.reel_strips[reel_idx] + self.reel_strips[reel_idx]
            drop_reel.append(long_reel[drop_pos[reel_idx]:drop_pos[reel_idx] + self.shape[reel_idx]])


        for x in range(len(pos)):
            reel_info[Const.R_Reel_Drop_Num][x] += self.shape[x]
        return {
            'drop_item_list': drop_reel,
            Const.R_Reel_Pos: pos,
            Const.R_Reel_Drop_Num: copy.deepcopy(reel_info[Const.R_Reel_Drop_Num])
        }


class StandardLineEvaluator(object):
    """
    处理reel结果，获得线上赢钱
    """
    def __init__(self, totalbet, spinreel, payline, paytable, betLine, wild_sub, linesym, wilds, wild):

        self.totalbet = totalbet  # 押注
        self.spinreel = spinreel  # 盘面结果

        self.paytable = paytable
        self.payline = payline  # 赢钱线list

        self.wildsub = wild_sub  # wild可替代图标数组

        self.line_sym = linesym  # 连线图标
        self.wilds = wilds  # wild类图标
        self.wild = wild  # wild
        self.betLine = betLine  # 押注线数

        self.linebet = totalbet / betLine  # line bet每条线基础押注

    def GetAllLine(self):
        line_id = 0
        line_result = []
        for line in self.payline:
            col = 0
            line_combo = []
            line_pos = []

            for row in line:
                line_combo.append(self.spinreel[col][row])
                line_pos.append([col, row])
                col += 1
            line_result.append({Const.R_Line_Id: line_id, Const.R_Line_Combo: line_combo, Const.R_Line_Pos: line_pos})
            line_id += 1
        return line_result

    def evaluateLine(self, oneline):
        line_combo = oneline[Const.R_Line_Combo]
        sym_long = 0
        wild_long = 0
        kind = line_combo[0]
        line_mul = 0
        long = 0

        if line_combo[0] in self.line_sym:
            if line_combo[0] in self.wildsub:
                for x in range(len(line_combo)):
                    if line_combo[x] == line_combo[0] or line_combo[x] in self.wilds:
                        sym_long += 1
                    else:
                        break

                kind = line_combo[0]
                line_mul = self.paytable[kind][sym_long - 1]
                long = sym_long

            elif line_combo[0] in self.wilds:
                for x in range(len(line_combo)):
                    if line_combo[x] in self.wilds:
                        wild_long += 1
                    else:
                        break

                if wild_long < len(line_combo):

                    normal_sym = line_combo[wild_long]

                    if normal_sym in self.line_sym:
                        for x in range(len(line_combo)):
                            if line_combo[x] == line_combo[wild_long] or line_combo[x] in self.wilds:
                                sym_long += 1
                            else:
                                break

                        if self.paytable[normal_sym][sym_long - 1] >= self.paytable[self.wild][wild_long - 1]:
                            line_mul = self.paytable[normal_sym][sym_long - 1]
                            long = sym_long
                            kind = normal_sym
                        else:
                            line_mul = self.paytable[self.wild][wild_long - 1]
                            long = wild_long
                            kind = self.wild
                    else:
                        line_mul = self.paytable[self.wild][wild_long - 1]
                        long = wild_long
                        kind = self.wild
                else:
                    line_mul = self.paytable[self.wild][wild_long - 1]
                    long = wild_long
                    kind = self.wild

        oneline[Const.R_Line_Kind] = kind
        oneline[Const.R_Line_Mul] = line_mul
        oneline[Const.R_Line_Win] = line_mul * self.linebet
        oneline[Const.R_Line_Pos] = oneline[Const.R_Line_Pos][:long]
        oneline[Const.R_Line_Long] = long

        return oneline

    def evaluate(self):
        line_result = {}
        line_ifo = self.GetAllLine()
        line_result[Const.R_Line] = []
        win_pos_list = []
        win_amount = 0
        wild_2_active = 0
        for simple_line in line_ifo:
            oneline = self.evaluateLine(simple_line)

            if oneline[Const.R_Line_Win] > 0:
                long = oneline[Const.R_Line_Long]

                wild_mul = 0
                for i in range(long):
                    if oneline[Const.R_Line_Combo][i] in [12, 13, 15]:
                        wild_mul += oneline[Const.R_Line_Combo][i] - 10
                        wild_2_active = 1

                if wild_mul != 0:
                    oneline[Const.R_Line_Win] = oneline[Const.R_Line_Win] * wild_mul

                win_amount += oneline[Const.R_Line_Win]
                line_result[Const.R_Line].append(oneline)

                for pos in oneline[Const.R_Line_Pos]:
                    if pos not in win_pos_list:
                        win_pos_list.append(pos)

        line_result[Const.R_Line_WinAmount] = win_amount
        line_result[Const.R_Win_Amount] = win_amount
        line_result[Const.R_Win_Pos_List] = win_pos_list
        line_result['wild_2_active'] = wild_2_active
        return line_result

    def free_evaluate(self, toad_wild):
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

                wild_mul = 0
                for i in range(long):
                    if 12 <= oneline[Const.R_Line_Combo][i] < 20:
                        wild_mul += oneline[Const.R_Line_Combo][i] - 10
                        wild_2_active = True

                if toad_wild in oneline[Const.R_Line_Combo][:long]:
                    wild_mul += toad_wild - 20

                if wild_mul != 0:
                    oneline[Const.R_Line_Win] = oneline[Const.R_Line_Win] * wild_mul

                win_amount += oneline[Const.R_Line_Win]
                line_result[Const.R_Line].append(oneline)

                for pos in oneline[Const.R_Line_Pos]:
                    if pos not in win_pos_list:
                        win_pos_list.append(pos)

        line_result[Const.R_Line_WinAmount] = win_amount
        line_result[Const.R_Win_Amount] = win_amount
        line_result[Const.R_Win_Pos_List] = win_pos_list
        line_result['wild_2_active'] = wild_2_active
        return line_result


def rand_list(one_list):
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
