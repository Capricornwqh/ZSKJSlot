from Toad_Config import *
import random
import Const as Const
import Slot as Slot
import copy




Game_Set = Const.C_Game_Set


class GameSlot(object):

    def ng_spin(self, totalbet, Reel_Choose, wild_set, buy_free=0, buy_super=0):

        c_shape = [6, 6, 6, 6, 6, 6]
        free_spin = 0
        free_trigger = False

        toad_pos = []
        coin_wild_pos_list = []
        coin_wild_mul = []
        wild_buff_pos = []

        reel_idx = Slot.rand_list(Reel_Choose)
        reel_info = Slot.GetReel(Base_ReelSets[reel_idx], c_shape).get_reel()
        item_list = reel_info['item_list']
        near_miss = 0

        # 第一次判断赢钱 和 free 触发，决定后续是否触发 天降横财
        win_info = Slot.StandardLineEvaluator(totalbet, item_list, C_PayLine, C_Paytable, C_BetLine, C_Wild_Sub, C_LineSym, Wilds, Wild).evaluate()


        if win_info[Const.R_Win_Amount] == 0 and Wild not in reel_info['item_list'][0]:
            # 天降横财玩法
            item_list, wild_buff_pos = wild_buff(item_list, wild_set, is_fg=0)
            reel_info['item_list'] = item_list

        # 给2-6列的 每个 wild（铜钱） 判断是否变成 2X、 3X、 5X Wild
        # base 使用不同的轴，wild 带倍数的概率不一样
        for x in range(6):
            for y in range(6):
                if item_list[x][y] == Wild_S:
                    if random.random() < wild_set['mul_prob'][reel_idx]:
                        item_list[x][y] = 10 + Slot.rand_list(wild_set['wild_mul_weight'])


        reel_info['item_list'] = item_list
        # bet 20 win 108

        win_info = Slot.StandardLineEvaluator(totalbet, item_list, C_PayLine, C_Paytable, C_BetLine, C_Wild_Sub, C_LineSym, Wilds, Wild).evaluate()
        total_win = win_info[Const.R_Win_Amount]

        if Wild in item_list[0]:
            near_miss = 1
            # 记录蟾蜍位置
            for i in range(6):
                if item_list[0][i] == Wild:
                    toad_pos.append([0, i])

            # 记录铜钱位置
            for x in [1, 2, 3, 4, 5]:
                for y in range(6):
                    if item_list[x][y] in [Wild_S, Wild_S2, Wild_S3, Wild_S5]:
                        coin_wild_pos_list.append([x, y])

                        if item_list[x][y] in [Wild_S2, Wild_S3, Wild_S5]:
                            mul = item_list[x][y] - 10
                            coin_wild_mul.append([[x, y], mul])
                        else:
                            coin_wild_mul.append([[x, y], 0])

            if len(coin_wild_pos_list) > 0:
                free_spin = 5


        ret = {
            'reel_info': reel_info,
            'wild_buff_pos': wild_buff_pos,
            'win_info': win_info,
            'win_amount': total_win,
            'free_spin': free_spin,
            'coin_wild_pos': coin_wild_pos_list,
            'coin_wild_mul': coin_wild_mul,
            'near_miss': near_miss,
            'toad_info': {
                'toad_pos': toad_pos,
                'coin_wild_pos': coin_wild_pos_list,
                'toad_mul': 0,
                'coin_wild_mul': coin_wild_mul,
                'collect_wild': 0,
                'toad_size': 1
            },
        }

        return ret


    def fg_spin(self, totalbet, wild_set, toad_info, free_spins, win_buff, free_total_win, is_fg=1, buy_free=0, buy_super=0,test_pos=[]):
        c_shape = [6, 6, 6, 6, 6, 6]
        trigger_free_spin = 0
        toad_pos = toad_info['toad_pos']
        coin_wild_pos_list = []
        coin_wild_mul = []

        wild_buff_pos = []
        wild_buff_pos_mul = []

        # reset free spin
        toad_info['free_spin'] = 0
        toad_size = toad_info['toad_size']

        # 按照 蟾蜍大小选轴
        reel_idx = Slot.rand_list(wild_set['strip_choose'][toad_size - 1])


        # 概率使用 free 4
        if win_buff == 1:
            reel_idx = 3

        if buy_free == 1 or buy_super == 1:
            reel_info = Slot.GetReel(Free_ReelSets[reel_idx], c_shape).get_reel(test_pos)
        else:
            reel_info = Slot.GetReel(Free_ReelSets[reel_idx], c_shape).get_reel(test_pos)

        item_list = reel_info['item_list']


        toad_wild = 20 + toad_info['toad_mul']
        # toad_wild = 20
        # 蟾蜍图标，对盘面进行覆盖 20+ 为蟾蜍wild 大于20的值表示蟾蜍的倍数
        for pos in toad_pos:
            try:
                item_list[pos[0]][pos[1]] = toad_wild
            except IndexError:
                print(pos)
                print('toad pos error')

        reel_info['item_list'] = item_list


        # 给1-6列的 每个 wild（铜钱） 判断是否变成 2X、 3X、 5X Wild
        for x in range(6):
            for y in range(6):
                if item_list[x][y] == Wild_S:
                    if random.random() < wild_set['mul_prob']:
                        item_list[x][y] = 10 + Slot.rand_list(wild_set['wild_mul_weight'])

        # 记录铜钱位置
        for x in range(6):
            for y in range(6):
                if item_list[x][y] in [Wild_S, Wild_S2, Wild_S3, Wild_S5]:
                    coin_wild_pos_list.append([x, y])

                    if item_list[x][y] in [Wild_S2, Wild_S3, Wild_S5]:
                        mul = item_list[x][y] - 10
                        coin_wild_mul.append([[x, y], mul])
                    else:
                        coin_wild_mul.append([[x, y], 0])


        free_wilds = copy.deepcopy(Wilds)
        free_wilds.append(toad_wild)
        free_line_sym = copy.deepcopy(C_LineSym)
        free_line_sym.append(toad_wild)

        win_info = Slot.StandardLineEvaluator(totalbet, item_list, C_PayLine, C_Paytable, C_BetLine, C_Wild_Sub, free_line_sym, free_wilds, Wild).free_evaluate(toad_wild)
        total_win = win_info[Const.R_Win_Amount]


        toad_info['coin_wild_pos'] = copy.deepcopy(coin_wild_pos_list)
        toad_info['coin_wild_mul'] = copy.deepcopy(coin_wild_mul)


        # 只有在有wild 的情况下，蟾蜍才会移动
        if len(coin_wild_pos_list) > 0:
            toad_info = toad_move_upgrade(toad_info)
            trigger_free_spin = toad_info['free_spin']

        new_toad_pos = copy.deepcopy(toad_info['toad_pos'])
        toad_size = toad_info['toad_size']

        fg_wild_num = 0
        # 先检查 蟾蜍升级 所需的铜钱wild 是否 不超过 2 个, 如果超过 2 个，则不可能进入天降横财的触发判断
        # free spin 为 0 且 按照蟾蜍的变长取触发的概率
        # wild 不会落在 蟾蜍 所在的区域
        # 天降横财 上限为 2 个铜钱
        fg_wild_num = [5, 9, 13, 16, 19, 0][toad_size - 1] - toad_info['collect_wild']

        if fg_wild_num <= 2 and free_spins == 0 and trigger_free_spin == 0:
            if random.random() < wild_set['fg_end_buff_prob'][toad_size-1]:

                if fg_wild_num <= 0:
                    print('false fg_wild_num')

                pos_record = {}
                for x in range(6):
                    for y in range(6):
                        if [x, y] not in new_toad_pos:
                            if x not in pos_record.keys():
                                pos_record[x] = [[x,y]]
                            else:
                                pos_record[x].append([x, y])

                for k,v in pos_record.items():
                    if len(v) > 1:
                        pos_record[k] = [random.choice(v)]

                choose_reel_idx = random.sample(list(pos_record.keys()), fg_wild_num)

                for reel_idx in choose_reel_idx:
                    pos = pos_record[reel_idx][0]
                    wild_buff_pos.append(pos)

                    if random.random() < wild_set['mul_prob']:
                        mul = Slot.rand_list(wild_set['wild_mul_weight'])
                        wild_buff_pos_mul.append([pos, mul])
                    else:
                        wild_buff_pos_mul.append([pos, 0])

                toad_info['coin_wild_pos'] = copy.deepcopy(wild_buff_pos)
                toad_info['coin_wild_mul'] = copy.deepcopy(wild_buff_pos_mul)

                new_toad_info = toad_move_upgrade(toad_info)
                trigger_free_spin = toad_info['free_spin']
                # print('天降触发')
                # print(wild_buff_pos)
                # print(wild_buff_pos_mul)

        if len(wild_buff_pos) > 0:
            toad_info = new_toad_info
            trigger_free_spin = toad_info['free_spin']

        ret = {
            'reel_info': reel_info,
            'wild_buff_pos': wild_buff_pos,
            'wild_buff_pos_mul': wild_buff_pos_mul,
            'win_info': win_info,
            'win_amount': total_win,
            'free_spin': trigger_free_spin,
            'coin_wild_pos': coin_wild_pos_list,
            'coin_wild_mul': coin_wild_mul,
            'toad_info': toad_info,
        }

        return ret


    def buy_ng_spin(self, totalbet, Reel_Choose, wild_set, buy_free=0, buy_super=0, test_pos=[]):

        c_shape = [6, 6, 6, 6, 6, 6]
        free_spin = 0
        free_trigger = False

        toad_pos = []
        coin_wild_pos_list = []
        coin_wild_mul = []
        wild_buff_pos = []

        if buy_free == 1 and buy_super == 0:
            reel_idx = 0
        elif buy_free == 0 and buy_super == 1:
            reel_idx = 1
        else:
            reel_idx = 0
            print('reel_idx false')

        reel_info = Slot.GetReel(Buy_Base_ReelSets[reel_idx], c_shape).get_reel(test_pos)

        if buy_super == 0:
            r2_r6 = copy.deepcopy(reel_info['item_list'][1:])
            random.shuffle(r2_r6)
            reel_info['item_list'] = [reel_info['item_list'][0]] + r2_r6


        item_list = reel_info['item_list']

        # buy super 的铜钱 wild，选一个赋值倍数
        if buy_super == 1:
            wild_choose_list = []
            for x in range(6):
                for y in range(6):
                    if item_list[x][y] == Wild_S:
                        wild_choose_list.append([x, y])

            if len(wild_choose_list) < 5:
                print('Buy Super Wild Less')

            choose_pos = random.choice(wild_choose_list)
            item_list[choose_pos[0]][choose_pos[1]] = 10 + Slot.rand_list(wild_set['wild_mul_weight'])

        # 给2-6列的 每个 wild（铜钱） 判断是否变成 2X、 3X、 5X Wild
        # base 使用不同的轴，wild 带倍数的概率不一样
        for x in range(6):
            for y in range(6):
                if item_list[x][y] == Wild_S:

                    if random.random() < wild_set['mul_prob'][reel_idx]:
                        item_list[x][y] = 10 + Slot.rand_list(wild_set['wild_mul_weight'])

        # a_item_list = [
        #     [LK, H3, LA, H2, H4, H1],
        #     [LK, H3, LA, LJ, 11, H2],
        #     [10, LQ, LQ, LJ, LJ, LQ],
        #     [LQ, LQ, LQ, H3, LJ, LQ],
        #     [LQ, H1, 11, 11, LJ, LJ],
        #     [LQ, H1, LK, LA, H2, LJ]
        # ]
        #
        # n_item_list = [[] for _ in range(6)]
        # for x in range(6):
        #     for y in range(6):
        #         n_item_list[y].append(a_item_list[x][y])


        reel_info['item_list'] = item_list
        # bet 20 win 108

        win_info = Slot.StandardLineEvaluator(totalbet, item_list, C_PayLine, C_Paytable, C_BetLine, C_Wild_Sub, C_LineSym, Wilds, Wild).evaluate()
        total_win = win_info[Const.R_Win_Amount]

        if Wild in item_list[0]:
            # 记录蟾蜍位置
            for i in range(6):
                if item_list[0][i] == Wild:
                    toad_pos.append([0, i])

            # 记录铜钱位置
            for x in [1, 2, 3, 4, 5]:
                for y in range(6):
                    if item_list[x][y] in [Wild_S, Wild_S2, Wild_S3, Wild_S5]:
                        coin_wild_pos_list.append([x, y])

                        if item_list[x][y] in [Wild_S2, Wild_S3, Wild_S5]:
                            mul = item_list[x][y] - 10
                            coin_wild_mul.append([[x, y], mul])
                        else:
                            coin_wild_mul.append([[x, y], 0])

            if len(coin_wild_pos_list) > 0:
                free_spin = 5


        ret = {
            'reel_info': reel_info,
            'wild_buff_pos': wild_buff_pos,
            'win_info': win_info,
            'win_amount': total_win,
            'free_spin': free_spin,
            'coin_wild_pos': coin_wild_pos_list,
            'coin_wild_mul': coin_wild_mul,

            'toad_info': {
                'toad_pos': toad_pos,
                'coin_wild_pos': coin_wild_pos_list,
                'toad_mul': 0,
                'coin_wild_mul': coin_wild_mul,
                'collect_wild': 0,
                'toad_size': 1
            },
        }

        return ret




def wild_buff(item_list, wild_set, is_fg):
    buff_pos = []
    if is_fg == 0:
        # base
        add_wild_num = 0
        if random.random() < wild_set['extra_wild_prob']:
            add_wild_num = Slot.rand_list(wild_set['extra_wild_num_weight'])

        if add_wild_num > 0:
            wild_reel_mark = [0 for _ in range(6)]
            # print('item_list_1',item_list)
            # print('add_wild_num',add_wild_num)
            # 随机选一条中奖线
            choose_line = random.choice(C_PayLine)
            # print('choose_line',choose_line)
            for x in [1, 2, 3, 4, 5]:
                # 把 所在列没有 wild 的列 记录下来
                if Wild_S not in item_list[x]:
                    wild_reel_mark[x] = 1
            # print('wild_reel_mark',wild_reel_mark)
            choose_reel = []
            for x in [1, 2, 3, 4, 5]:
                if wild_reel_mark[x] == 1:
                    choose_reel.append(x)

            if len(choose_reel) < add_wild_num:
                add_wild_num = len(choose_reel)

            # print('choose_reel',choose_reel)
            # 按顺序从 2 - 6 截取对应数量的 不含 wild 的列
            r_reels = choose_reel[:add_wild_num]
            # 按照 中奖线，替换成 铜钱wild
            for r_idx in r_reels:
                c_idx = choose_line[r_idx]
                item_list[r_idx][c_idx] = Wild_S
                buff_pos.append([r_idx, c_idx])
            # print('item_list_2',item_list)
    return item_list, buff_pos




def toad_move_upgrade(toad_info):
    toad_pos = toad_info['toad_pos']
    coin_wild_pos = copy.deepcopy(toad_info['coin_wild_pos'])
    toad_mul = toad_info['toad_mul']
    coin_wild_mul = toad_info['coin_wild_mul']
    collect_num = toad_info['collect_wild']
    toad_size = toad_info['toad_size']

    free_spin = 0

    while len(coin_wild_pos) > 0:
        eat_pos = []
        # print('toad_pos', toad_pos)
        landing_point = find_landing_point(toad_pos, coin_wild_pos)
        # print(landing_point)
        # 计算 蟾蜍 Wild 最终落点，的位移向量
        move_vector = cal_landing_vector(toad_pos, landing_point)
        # print('move_vector', move_vector)
        toad_pos = toad_move(toad_pos, move_vector)

        coin_wild_pos.remove(landing_point)
        eat_pos.append(landing_point)

        for pos in coin_wild_pos:
            if pos in toad_pos:
                coin_wild_pos.remove(pos)
                eat_pos.append(pos)

        # print('eat_pos', eat_pos)

        collect_num += len(eat_pos)

    # 蟾蜍移动之后，才会增加倍数
    for wild_mul in coin_wild_mul:
        if wild_mul[1] > 0:
            toad_mul += wild_mul[1]


    # 可能会出现连续升级的情况（不排除）
    if collect_num >= 5 and toad_size == 1:
        free_spin += 3
        toad_pos = toad_expand(toad_pos, toad_size)
        toad_size += 1

    if collect_num >= 9 and toad_size == 2:
        free_spin += 3
        toad_pos = toad_expand(toad_pos, toad_size)
        toad_size += 1

    if collect_num >= 13 and toad_size == 3:
        free_spin += 2
        toad_pos = toad_expand(toad_pos, toad_size)
        toad_size += 1

    if collect_num >= 16 and toad_size == 4:
        free_spin += 2
        toad_pos = toad_expand(toad_pos, toad_size)
        toad_size += 1

    if collect_num >= 19 and toad_size == 5:
        free_spin += 1
        toad_pos = toad_expand(toad_pos, toad_size)
        toad_size += 1

    coin_wild_mul = []
    coin_wild_pos = []

    ret = {
        'toad_pos': toad_pos,
        'toad_size': toad_size,
        'coin_wild_pos': coin_wild_pos,
        'toad_mul': toad_mul,
        'coin_wild_mul': coin_wild_mul,
        'collect_wild': collect_num,
        'free_spin': free_spin,
    }
    return ret


# 计算蟾蜍的落地位置
def find_landing_point(toad_pos, coin_wild_pos_list):

    # 获取所有 铜钱 wild 到蟾蜍 wild的距离
    distance_dict = {}
    for coin_wild_pos in coin_wild_pos_list:
        distance = cal_distance(toad_pos, coin_wild_pos)

        if distance in distance_dict.keys():
            distance_dict[distance].append(coin_wild_pos)
        else:
            distance_dict[distance] = [coin_wild_pos]

    # 找出距离最小的点
    min_distance = min(distance_dict.keys())

    if len(distance_dict[min_distance]) > 1:
        # 找出金蟾的 4/1 个顶点
        x_idx = []
        y_idx = []
        for pos in toad_pos:
            x_idx.append(pos[0])
            y_idx.append(pos[1])

        max_x = max(x_idx)
        min_x = min(x_idx)
        max_y = max(y_idx)
        min_y = min(y_idx)

        area_dict = {1: [], 2: [], 3: [], 4: [], 5: [], 6: [], 7: [], 8: []}

        has_pos_area = []
        for pos in distance_dict[min_distance]:
            m = pos[0]
            n = pos[1]

            if max_x < m and min_y <= n <= max_y:
                area_dict[1].append(pos)
                has_pos_area.append(1)
            elif max_x < m and n < min_y:
                area_dict[2].append(pos)
                has_pos_area.append(2)
            elif max_x < m and max_y < n:
                area_dict[3].append(pos)
                has_pos_area.append(3)
            elif m < min_x and min_y <= n <= max_y:
                area_dict[4].append(pos)
                has_pos_area.append(4)
            elif m < min_x and n < min_y:
                area_dict[5].append(pos)
                has_pos_area.append(5)
            elif m < min_x and max_y < n:
                area_dict[6].append(pos)
                has_pos_area.append(6)
            elif min_x <= m <= max_x and n < min_y:
                area_dict[7].append(pos)
                has_pos_area.append(7)
            elif min_x <= m <= max_x and max_y < n:
                area_dict[8].append(pos)
                has_pos_area.append(8)
            else:
                print('find_landing_point error')

        advance_area = min(has_pos_area)

        if len(area_dict[advance_area]) > 1:

            points = area_dict[advance_area]
            max_x_2 = max(p[0] for p in points)  # 找到最大x值
            candidates = [p for p in points if p[0] == max_x_2]  # 筛选x最大的点
            landing_point = min(candidates, key=lambda p: p[1])  # 在候选点中找y最小的
            return landing_point

        else:
            landing_point = area_dict[advance_area][0]

    else:
        landing_point = distance_dict[min_distance][0]

    return landing_point


# 获取 铜钱wild，和 (蟾蜍 wild / 指定坐标)  的距离
def cal_distance(toad_pos, c_pos):
    min_distance = float('inf')

    # if len(toad_pos) > 4:
    #     square_pos = get_square_vertices(toad_pos)
    # else:
    #     square_pos = toad_pos

    square_pos = toad_pos

    for t_pos in square_pos:
        distance = abs(t_pos[0] - c_pos[0]) + abs(t_pos[1] - c_pos[1])

        if distance < min_distance:
            min_distance = distance

    return min_distance


def get_square_vertices(coordinates):
    # 提取所有x和y坐标
    xs = [point[0] for point in coordinates]
    ys = [point[1] for point in coordinates]

    # 计算x和y的最小、最大值
    min_x, max_x = min(xs), max(xs)
    min_y, max_y = min(ys), max(ys)

    # 生成四个顶点
    vertices = [
        [min_x, min_y],
        [min_x, max_y],
        [max_x, min_y],
        [max_x, max_y]
    ]

    return vertices


# 计算 蟾蜍 Wild 最终落点，的位移向量
def cal_landing_vector(toad_pos, landing_point):

    # 找到蟾蜍 wild 距离落点最近的点

    min_distance = float('inf')
    for t_pos in toad_pos:
        distance = abs(t_pos[0] - landing_point[0]) + abs(t_pos[1] - landing_point[1])
        if distance < min_distance:
            min_distance = distance
            toad_landing_pos = t_pos

    # 获得向量值
    vector = [landing_point[0] - toad_landing_pos[0], landing_point[1] - toad_landing_pos[1]]
    return vector


# 蟾蜍 wild 按照向量方向移动
def toad_move(toad_pos, vector):
    new_toad_pos = []
    for t_pos in toad_pos:
        new_pos = [t_pos[0] + vector[0], t_pos[1] + vector[1]]
        new_toad_pos.append(new_pos)
    return new_toad_pos


# 蟾蜍变大，直接穷举了
def toad_expand(toad_pos, Size):
    x_list = []
    y_list = []
    mark_pos = []
    for pos in toad_pos:
        if pos[0] in x_list:
            pass
        else:
            x_list.append(pos[0])
        if pos[1] in y_list:
            pass
        else:
            y_list.append(pos[1])
    mark_pos.append(min(x_list))
    mark_pos.append(max(y_list))
    if Size == 1:
        if mark_pos[0] <= 2 and mark_pos[1] <= 2:
            toad_pos.append([mark_pos[0] + 1, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0], mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1]])
        elif mark_pos[0] >= 3 and mark_pos[1] <= 2:
            toad_pos.append([mark_pos[0] - 1, mark_pos[1]])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0], mark_pos[1] + 1])
        elif mark_pos[0] <= 2 and mark_pos[1] >= 3:
            toad_pos.append([mark_pos[0], mark_pos[1] - 1])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1]])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1] - 1])
        elif mark_pos[0] >= 3 and mark_pos[1] >= 3:
            toad_pos.append([mark_pos[0] - 1, mark_pos[1]])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 1])
            toad_pos.append([mark_pos[0], mark_pos[1] - 1])
    elif Size == 2:
        if mark_pos[0] <= 2 and mark_pos[1] <= 2:
            toad_pos.append([mark_pos[0], mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 2, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 2, mark_pos[1]])
            toad_pos.append([mark_pos[0] + 2, mark_pos[1] - 1])
        elif mark_pos[0] >= 3 and mark_pos[1] <= 2:
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 1])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1]])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0], mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1] + 1])
        elif mark_pos[0] <= 2 and mark_pos[1] >= 3:
            toad_pos.append([mark_pos[0] + 2, mark_pos[1]])
            toad_pos.append([mark_pos[0] + 2, mark_pos[1] - 1])
            toad_pos.append([mark_pos[0] + 2, mark_pos[1] - 2])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1] - 2])
            toad_pos.append([mark_pos[0], mark_pos[1] - 2])
        elif mark_pos[0] >= 3 and mark_pos[1] >= 3:
            toad_pos.append([mark_pos[0] - 1, mark_pos[1]])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 1])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 2])
            toad_pos.append([mark_pos[0], mark_pos[1] - 2])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1] - 2])
    elif Size == 3:
        if mark_pos[0] <= 1 and mark_pos[1] <= 3:
            toad_pos.append([mark_pos[0], mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 2, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 3, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 3, mark_pos[1]])
            toad_pos.append([mark_pos[0] + 3, mark_pos[1] - 1])
            toad_pos.append([mark_pos[0] + 3, mark_pos[1] - 2])
        elif mark_pos[0] >= 2 and mark_pos[1] <= 3:
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 2])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 1])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1]])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0], mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 2, mark_pos[1] + 1])
        elif mark_pos[0] <= 1 and mark_pos[1] >= 4:
            toad_pos.append([mark_pos[0], mark_pos[1] - 3])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1] - 3])
            toad_pos.append([mark_pos[0] + 2, mark_pos[1] - 3])
            toad_pos.append([mark_pos[0] + 3, mark_pos[1] - 3])
            toad_pos.append([mark_pos[0] + 3, mark_pos[1] - 2])
            toad_pos.append([mark_pos[0] + 3, mark_pos[1] - 1])
            toad_pos.append([mark_pos[0] + 3, mark_pos[1]])
        elif mark_pos[0] >= 2 and mark_pos[1] >= 4:
            toad_pos.append([mark_pos[0] - 1, mark_pos[1]])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 1])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 2])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 3])
            toad_pos.append([mark_pos[0], mark_pos[1] - 3])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1] - 3])
            toad_pos.append([mark_pos[0] + 2, mark_pos[1] - 3])
    elif Size == 4:
        if mark_pos[0] <= 1 and mark_pos[1] <= 3:
            toad_pos.append([mark_pos[0], mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 2, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 3, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 4, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 4, mark_pos[1]])
            toad_pos.append([mark_pos[0] + 4, mark_pos[1] - 1])
            toad_pos.append([mark_pos[0] + 4, mark_pos[1] - 2])
            toad_pos.append([mark_pos[0] + 4, mark_pos[1] - 3])
        elif mark_pos[0] >= 2 and mark_pos[1] <= 3:
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 3])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 2])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 1])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1]])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 0, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 2, mark_pos[1] + 1])
            toad_pos.append([mark_pos[0] + 3, mark_pos[1] + 1])
        elif mark_pos[0] <= 1 and mark_pos[1] >= 4:
            toad_pos.append([mark_pos[0], mark_pos[1] - 4])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1] - 4])
            toad_pos.append([mark_pos[0] + 2, mark_pos[1] - 4])
            toad_pos.append([mark_pos[0] + 3, mark_pos[1] - 4])
            toad_pos.append([mark_pos[0] + 4, mark_pos[1] - 4])
            toad_pos.append([mark_pos[0] + 4, mark_pos[1]])
            toad_pos.append([mark_pos[0] + 4, mark_pos[1] - 1])
            toad_pos.append([mark_pos[0] + 4, mark_pos[1] - 2])
            toad_pos.append([mark_pos[0] + 4, mark_pos[1] - 3])
        elif mark_pos[0] >= 2 and mark_pos[1] >= 4:
            toad_pos.append([mark_pos[0] - 1, mark_pos[1]])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 1])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 2])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 3])
            toad_pos.append([mark_pos[0] - 1, mark_pos[1] - 4])
            toad_pos.append([mark_pos[0], mark_pos[1] - 4])
            toad_pos.append([mark_pos[0] + 1, mark_pos[1] - 4])
            toad_pos.append([mark_pos[0] + 2, mark_pos[1] - 4])
            toad_pos.append([mark_pos[0] + 3, mark_pos[1] - 4])
    elif Size == 5:
        for x in range(6):
            for y in range(6):
                if [x, y] not in toad_pos:
                    toad_pos.append([x, y])
                else:
                    pass

    return toad_pos
