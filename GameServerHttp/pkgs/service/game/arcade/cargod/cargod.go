package service_cargod

import (
	"math/rand"
)

type CarGodService struct {
	Car *Car // 车神游戏配置：包含轮盘位置、赔率、权重等所有游戏数据
}

// 创建新的车神游戏
func NewCarGodService() *CarGodService {
	// 初始化车神游戏服务
	// 这里可以加载配置文件或数据库数据，初始化Car结构体
	return &CarGodService{
		Car: &Car{
			CarType: []string{
				"Ferrari", "Ferrari", "Ferrari", "Benz", "Benz", "Benz", "BMW", "BMW", "BMW", "Audi", "Audi", "Audi",
				"Ferrari", "Ferrari", "Ferrari", "Benz", "Benz", "Benz", "BMW", "BMW", "BMW", "Audi", "Audi", "Audi",
			},
			CarColor: []string{
				"red", "green", "yellow", "red", "green", "yellow", "red", "green", "yellow", "red", "green", "yellow",
				"red", "green", "yellow", "red", "green", "yellow", "red", "green", "yellow", "red", "green", "yellow",
			},
			GroupSet: &GroupSet{
				SymbolPosition: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
				EventWeight:    []int{900, 5000, 8800, 5500, 1000, 8800, 7800, 2300, 1200, 500, 500, 500, 400, 500, 600, 700, 400, 400, 400, 0, 100, 1500, 0},
				JackpotOdds:    []int{60, 70, 80, 90, 100},
				JackpotWeight:  []int{35, 30, 20, 10, 5},
			},
			LevelSet: []*LevelSet{
				{
					Odds:           []int{41, 28, 20, 8, 6, 4, 19, 17, 13, 12, 10, 7},
					PositionWeight: []int{7368, 12917, 18860, 52820, 71969, 110912, 21184, 24123, 32581, 34857, 42725, 63825, 7368, 12917, 18860, 52820, 71969, 110912, 21184, 24123, 32581, 34857, 42725, 63825},
					SpecialRate:    []int{700, 420, 544, 681, 478, 900, 567, 476, 348, 379, 682, 505, 700, 420, 544, 681, 478, 900, 567, 476, 348, 379, 682, 505},
				},
				{
					Odds:           []int{41, 20, 28, 8, 4, 6, 19, 13, 17, 12, 7, 10},
					PositionWeight: []int{7278, 19820, 11177, 51750, 110482, 74069, 21184, 32711, 23843, 34857, 63375, 43515, 7278, 19820, 11177, 51750, 110482, 74069, 21184, 32711, 23843, 34857, 63375, 43515},
					SpecialRate:    []int{750, 544, 420, 681, 900, 478, 567, 348, 476, 379, 505, 682, 750, 544, 420, 681, 900, 478, 567, 348, 476, 379, 505, 682},
				},
				{
					Odds:           []int{28, 41, 20, 6, 8, 4, 17, 19, 13, 10, 12, 7},
					PositionWeight: []int{13417, 6908, 18860, 72409, 52190, 110912, 24123, 20984, 32311, 42975, 35057, 63945, 13417, 6908, 18860, 72409, 52190, 110912, 24123, 20984, 32311, 42975, 35057, 63945},
					SpecialRate:    []int{420, 800, 544, 478, 681, 900, 476, 567, 348, 682, 379, 505, 420, 800, 544, 478, 681, 900, 476, 567, 348, 682, 379, 505},
				},
				{
					Odds:           []int{28, 20, 41, 6, 4, 8, 17, 13, 19, 10, 7, 12},
					PositionWeight: []int{13332, 19760, 5238, 71639, 110422, 54220, 24128, 32586, 20364, 42725, 63825, 35902, 13332, 19760, 5238, 71639, 110422, 54220, 24128, 32586, 20364, 42725, 63825, 35902},
					SpecialRate:    []int{450, 544, 700, 478, 900, 721, 476, 348, 567, 682, 505, 379, 450, 544, 700, 478, 900, 721, 476, 348, 567, 682, 505, 379},
				},
				{
					Odds:           []int{20, 41, 28, 4, 8, 6, 13, 19, 17, 7, 12, 10},
					PositionWeight: []int{20260, 6188, 10057, 110092, 53870, 74359, 32811, 20944, 23483, 63555, 35027, 43485, 20260, 6188, 10057, 110092, 53870, 74359, 32811, 20944, 23483, 63555, 35027, 43485},
					SpecialRate:    []int{544, 700, 420, 900, 681, 478, 398, 567, 476, 505, 379, 682, 544, 700, 420, 900, 681, 478, 398, 567, 476, 505, 379, 682},
				},
				{
					Odds:           []int{20, 28, 41, 4, 6, 8, 13, 17, 19, 7, 10, 12},
					PositionWeight: []int{20260, 12267, 4018, 110022, 73609, 54665, 32796, 23883, 20284, 63265, 43305, 35767, 20260, 12267, 4018, 110022, 73609, 54665, 32796, 23883, 20284, 63265, 43305, 35767},
					SpecialRate:    []int{544, 420, 700, 900, 478, 681, 348, 476, 567, 505, 721, 379, 544, 420, 700, 900, 478, 681, 348, 476, 567, 505, 721, 379},
				},
				{
					Odds:           []int{43, 24, 21, 7, 6, 4, 17, 15, 14, 13, 10, 8},
					PositionWeight: []int{6578, 15502, 17305, 60450, 71219, 110212, 23804, 27478, 29536, 31407, 42585, 55225, 6578, 15502, 17305, 60450, 71219, 110212, 23804, 27478, 29536, 31407, 42585, 55225},
					SpecialRate:    []int{800, 433, 554, 681, 496, 910, 621, 496, 376, 430, 682, 530, 800, 433, 554, 681, 496, 910, 621, 496, 376, 430, 682, 530},
				},
				{
					Odds:           []int{43, 21, 24, 7, 4, 6, 17, 14, 15, 13, 8, 10},
					PositionWeight: []int{6538, 18305, 13702, 59520, 109522, 73679, 23834, 29746, 27298, 31487, 54505, 43165, 6538, 18305, 13702, 59520, 109522, 73679, 23834, 29746, 27298, 31487, 54505, 43165},
					SpecialRate:    []int{800, 554, 433, 681, 910, 496, 621, 376, 496, 430, 530, 682, 800, 554, 433, 681, 910, 496, 621, 376, 496, 430, 530, 682},
				},
				{
					Odds:           []int{24, 43, 21, 6, 7, 4, 15, 17, 14, 10, 13, 8},
					PositionWeight: []int{16067, 6158, 17365, 71739, 59790, 110172, 27478, 23614, 29241, 42735, 31757, 55185, 16067, 6158, 17365, 71739, 59790, 110172, 27478, 23614, 29241, 42735, 31757, 55185},
					SpecialRate:    []int{433, 800, 554, 496, 681, 910, 496, 621, 376, 682, 430, 530, 433, 800, 554, 496, 681, 910, 496, 621, 376, 682, 430, 530},
				},
				{
					Odds:           []int{24, 21, 43, 6, 4, 7, 15, 14, 17, 10, 8, 13},
					PositionWeight: []int{16007, 18460, 4468, 70779, 109217, 62370, 27584, 29556, 23074, 42450, 54925, 32417, 16007, 18460, 4468, 70779, 109217, 62370, 27584, 29556, 23074, 42450, 54925, 32417},
					SpecialRate:    []int{433, 554, 800, 496, 910, 681, 496, 376, 621, 682, 530, 430, 433, 554, 800, 496, 910, 681, 496, 376, 621, 682, 530, 430},
				},
				{
					Odds:           []int{21, 43, 24, 4, 7, 6, 14, 17, 15, 8, 13, 10},
					PositionWeight: []int{18980, 5438, 12667, 109492, 61964, 74359, 29922, 23729, 27208, 55025, 31942, 43405, 18980, 5438, 12667, 109492, 61964, 74359, 29922, 23729, 27208, 55025, 31942, 43405},
					SpecialRate:    []int{660, 700, 450, 910, 722, 478, 452, 567, 476, 505, 379, 682, 660, 700, 450, 910, 722, 478, 452, 567, 476, 505, 379, 682},
				},
				{
					Odds:           []int{21, 24, 43, 4, 6, 7, 14, 15, 17, 8, 10, 13},
					PositionWeight: []int{18835, 14932, 3278, 109012, 72959, 62460, 29786, 27288, 23094, 54325, 42925, 32407, 18835, 14932, 3278, 109012, 72959, 62460, 29786, 27288, 23094, 54325, 42925, 32407},
					SpecialRate:    []int{580, 433, 800, 910, 496, 681, 376, 496, 621, 530, 682, 430, 580, 433, 800, 910, 496, 681, 376, 496, 621, 530, 682, 430},
				},
				{
					Odds:           []int{45, 29, 19, 7, 6, 5, 17, 13, 12, 11, 9, 8},
					PositionWeight: []int{6228, 12117, 19794, 61300, 72380, 88527, 23874, 32623, 35381, 38277, 48085, 55555, 6228, 12117, 19794, 61300, 72380, 88527, 23874, 32623, 35381, 38277, 48085, 55555},
					SpecialRate:    []int{800, 460, 544, 741, 478, 900, 587, 522, 348, 450, 682, 552, 800, 460, 544, 741, 478, 900, 587, 522, 348, 450, 682, 552},
				},
				{
					Odds:           []int{45, 19, 29, 7, 5, 6, 17, 12, 13, 11, 8, 9},
					PositionWeight: []int{6123, 20944, 10552, 60560, 87957, 74445, 23964, 35501, 32253, 38307, 54740, 48795, 6123, 20944, 10552, 60560, 87957, 74445, 23964, 35501, 32253, 38307, 54740, 48795},
					SpecialRate:    []int{800, 544, 460, 741, 900, 478, 587, 348, 522, 450, 552, 682, 800, 544, 460, 741, 900, 478, 587, 348, 522, 450, 552, 682},
				},
				{
					Odds:           []int{29, 45, 19, 6, 7, 5, 13, 17, 12, 9, 11, 8},
					PositionWeight: []int{12767, 5737, 19960, 72549, 60685, 88612, 32613, 23754, 35091, 48085, 38687, 55551, 12767, 5737, 19960, 72549, 60685, 88612, 32613, 23754, 35091, 48085, 38687, 55551},
					SpecialRate:    []int{460, 800, 544, 478, 741, 900, 522, 587, 348, 682, 450, 552, 460, 800, 544, 478, 741, 900, 522, 587, 348, 682, 450, 552},
				},
				{
					Odds:           []int{29, 19, 45, 6, 5, 7, 13, 12, 17, 9, 8, 11},
					PositionWeight: []int{12117, 19794, 6228, 72380, 88527, 61300, 32623, 35381, 23874, 48085, 55555, 38277, 12117, 19794, 6228, 72380, 88527, 61300, 32623, 35381, 23874, 48085, 55555, 38277},
					SpecialRate:    []int{450, 544, 700, 478, 900, 721, 476, 348, 567, 682, 505, 379, 450, 544, 700, 478, 900, 721, 476, 348, 567, 682, 505, 379},
				},
				{
					Odds:           []int{19, 45, 29, 5, 7, 6, 12, 17, 13, 8, 11, 9},
					PositionWeight: []int{21490, 5336, 9957, 87397, 62164, 74348, 35701, 23650, 32078, 54842, 38505, 48673, 21490, 5336, 9957, 87397, 62164, 74348, 35701, 23650, 32078, 54842, 38505, 48673},
					SpecialRate:    []int{564, 800, 520, 900, 721, 508, 348, 622, 476, 562, 392, 695, 564, 800, 520, 900, 721, 508, 348, 622, 476, 562, 392, 695},
				},
				{
					Odds:           []int{19, 29, 45, 5, 6, 7, 12, 13, 17, 8, 9, 11},
					PositionWeight: []int{21494, 11917, 3568, 87607, 73370, 62750, 35656, 32398, 23154, 54535, 48380, 39312, 21494, 11917, 3568, 87607, 73370, 62750, 35656, 32398, 23154, 54535, 48380, 39312},
					SpecialRate:    []int{544, 460, 800, 900, 478, 741, 348, 522, 587, 552, 682, 450, 544, 460, 800, 900, 478, 741, 348, 522, 587, 552, 682, 450},
				},
				{
					Odds:           []int{62, 51, 40, 6, 5, 4, 33, 25, 22, 12, 9, 7},
					PositionWeight: []int{4358, 5837, 7254, 70710, 86170, 109447, 11044, 15368, 17481, 34157, 46745, 61960, 4358, 5837, 7254, 70710, 86170, 109447, 11044, 15368, 17481, 34157, 46745, 61960},
					SpecialRate:    []int{344, 360, 544, 553, 378, 430, 386, 422, 348, 450, 482, 552, 344, 360, 544, 553, 378, 430, 386, 422, 348, 450, 482, 552},
				},
				{
					Odds:           []int{62, 40, 51, 6, 4, 5, 33, 22, 25, 12, 7, 9},
					PositionWeight: []int{4318, 8374, 4617, 70530, 108527, 87380, 11059, 17856, 14828, 33957, 61330, 47775, 4318, 8374, 4617, 70530, 108527, 87380, 11059, 17856, 14828, 33957, 61330, 47775},
					SpecialRate:    []int{344, 544, 360, 553, 430, 378, 386, 348, 422, 450, 552, 482, 344, 544, 360, 553, 430, 378, 386, 348, 422, 450, 552, 482},
				},
				{
					Odds:           []int{51, 62, 40, 5, 6, 4, 25, 33, 22, 9, 12, 7},
					PositionWeight: []int{6397, 3898, 7384, 85970, 70710, 109447, 15568, 10844, 17341, 46905, 34157, 61960, 6397, 3898, 7384, 85970, 70710, 109447, 15568, 10844, 17341, 46905, 34157, 61960},
					SpecialRate:    []int{360, 344, 544, 378, 553, 430, 422, 386, 348, 482, 450, 552, 360, 344, 544, 378, 553, 430, 422, 386, 348, 482, 450, 552},
				},
				{
					Odds:           []int{51, 40, 62, 5, 4, 6, 25, 22, 33, 9, 7, 12},
					PositionWeight: []int{6327, 8454, 2798, 85810, 108147, 72140, 15518, 17851, 10144, 46625, 61520, 35197, 6327, 8454, 2798, 85810, 108147, 72140, 15518, 17851, 10144, 46625, 61520, 35197},
					SpecialRate:    []int{360, 544, 344, 378, 430, 553, 422, 348, 386, 482, 552, 450, 360, 544, 344, 378, 430, 553, 422, 348, 386, 482, 552, 450},
				},
				{
					Odds:           []int{40, 62, 51, 4, 6, 5, 22, 33, 25, 7, 12, 9},
					PositionWeight: []int{8884, 3798, 4477, 108067, 71235, 87425, 18081, 10694, 14468, 61280, 34417, 47705, 8884, 3798, 4477, 108067, 71235, 87425, 18081, 10694, 14468, 61280, 34417, 47705},
					SpecialRate:    []int{544, 344, 360, 430, 553, 378, 348, 386, 422, 552, 450, 482, 544, 344, 360, 430, 553, 378, 348, 386, 422, 552, 450, 482},
				},
				{
					Odds:           []int{40, 51, 62, 4, 5, 6, 22, 25, 33, 7, 9, 12},
					PositionWeight: []int{8874, 5837, 2543, 108172, 86280, 72070, 18001, 15198, 9944, 61160, 47345, 35107, 8874, 5837, 2543, 108172, 86280, 72070, 18001, 15198, 9944, 61160, 47345, 35107},
					SpecialRate:    []int{544, 360, 344, 430, 378, 553, 348, 422, 386, 552, 482, 450, 544, 360, 344, 430, 378, 553, 348, 422, 386, 552, 482, 450},
				},
				{
					Odds:           []int{70, 62, 44, 6, 5, 3, 46, 38, 24, 16, 10, 8},
					PositionWeight: []int{3613, 4362, 6309, 70715, 85875, 146347, 7284, 9243, 15971, 24882, 42000, 54130, 3613, 4362, 6309, 70715, 85875, 146347, 7284, 9243, 15971, 24882, 42000, 54130},
					SpecialRate:    []int{344, 360, 544, 522, 378, 418, 386, 422, 348, 450, 482, 552, 344, 360, 544, 522, 378, 418, 386, 422, 348, 450, 482, 552},
				},
				{
					Odds:           []int{70, 44, 62, 6, 3, 5, 46, 24, 38, 16, 8, 10},
					PositionWeight: []int{3598, 7449, 3032, 70305, 145307, 87420, 7264, 16371, 8843, 24772, 53440, 42930, 3598, 7449, 3032, 70305, 145307, 87420, 7264, 16371, 8843, 24772, 53440, 42930},
					SpecialRate:    []int{344, 544, 360, 522, 395, 378, 386, 380, 437, 450, 552, 482, 344, 544, 360, 522, 395, 378, 386, 380, 437, 450, 552, 482},
				},
				{
					Odds:           []int{62, 70, 44, 5, 6, 3, 38, 46, 24, 10, 16, 8},
					PositionWeight: []int{4882, 3143, 6409, 85875, 70715, 146347, 9363, 7024, 15771, 42180, 24892, 54130, 4882, 3143, 6409, 85875, 70715, 146347, 9363, 7024, 15771, 42180, 24892, 54130},
					SpecialRate:    []int{360, 344, 544, 378, 522, 418, 422, 386, 348, 482, 450, 552, 360, 344, 544, 378, 522, 418, 422, 386, 348, 482, 450, 552},
				},
				{
					Odds:           []int{62, 44, 70, 5, 3, 6, 38, 24, 46, 10, 8, 16},
					PositionWeight: []int{4832, 7484, 1848, 85475, 145137, 72365, 9358, 16291, 6384, 41900, 53820, 25837, 4832, 7484, 1848, 85475, 145137, 72365, 9358, 16291, 6384, 41900, 53820, 25837},
					SpecialRate:    []int{360, 544, 344, 378, 418, 522, 422, 348, 386, 482, 552, 450, 360, 544, 344, 378, 418, 522, 422, 348, 386, 482, 552, 450},
				},
				{
					Odds:           []int{44, 70, 62, 3, 6, 5, 24, 46, 38, 8, 16, 10},
					PositionWeight: []int{7919, 2853, 2522, 144947, 71515, 87455, 16491, 6984, 8483, 53540, 25202, 42820, 7919, 2853, 2522, 144947, 71515, 87455, 16491, 6984, 8483, 53540, 25202, 42820},
					SpecialRate:    []int{544, 344, 360, 418, 522, 378, 348, 386, 422, 552, 450, 482, 544, 344, 360, 418, 522, 378, 348, 386, 422, 552, 450, 482},
				},
				{
					Odds:           []int{44, 62, 70, 3, 5, 6, 24, 38, 46, 8, 10, 16},
					PositionWeight: []int{7809, 4112, 1313, 144947, 86525, 72395, 16491, 9083, 6264, 53430, 42540, 25822, 7809, 4112, 1313, 144947, 86525, 72395, 16491, 9083, 6264, 53430, 42540, 25822},
					SpecialRate:    []int{544, 360, 344, 418, 378, 522, 348, 422, 386, 552, 482, 450, 544, 360, 344, 418, 378, 522, 348, 422, 386, 552, 482, 450},
				},
				{
					Odds:           []int{75, 65, 55, 7, 4, 3, 45, 35, 25, 15, 11, 9},
					PositionWeight: []int{3193, 4102, 4289, 59715, 107865, 145947, 7484, 10283, 15211, 26922, 38000, 47720, 3193, 4102, 4289, 59715, 107865, 145947, 7484, 10283, 15211, 26922, 38000, 47720},
					SpecialRate:    []int{344, 360, 544, 522, 378, 452, 386, 422, 348, 450, 482, 552, 344, 360, 544, 522, 378, 452, 386, 422, 348, 450, 482, 552},
				},
				{
					Odds:           []int{75, 55, 65, 7, 3, 4, 45, 25, 35, 15, 9, 11},
					PositionWeight: []int{3158, 5384, 2712, 59215, 144717, 109725, 7464, 15541, 9863, 26782, 47360, 38810, 3158, 5384, 2712, 59215, 144717, 109725, 7464, 15541, 9863, 26782, 47360, 38810},
					SpecialRate:    []int{344, 544, 360, 522, 452, 378, 386, 348, 422, 450, 552, 482, 344, 544, 360, 522, 452, 378, 386, 348, 422, 450, 552, 482},
				},
				{
					Odds:           []int{65, 75, 55, 4, 7, 3, 35, 45, 25, 11, 15, 9},
					PositionWeight: []int{4532, 2643, 4109, 108065, 59715, 145947, 10423, 7284, 15101, 38110, 26982, 47820, 4532, 2643, 4109, 108065, 59715, 145947, 10423, 7284, 15101, 38110, 26982, 47820},
					SpecialRate:    []int{360, 344, 544, 378, 522, 452, 422, 386, 348, 482, 450, 552, 360, 344, 544, 378, 522, 452, 422, 386, 348, 482, 450, 552},
				},
				{
					Odds:           []int{65, 55, 75, 4, 3, 7, 35, 25, 45, 11, 9, 15},
					PositionWeight: []int{4492, 5259, 1073, 107545, 144947, 61625, 10383, 15491, 6744, 37850, 47550, 27772, 4492, 5259, 1073, 107545, 144947, 61625, 10383, 15491, 6744, 37850, 47550, 27772},
					SpecialRate:    []int{360, 544, 344, 378, 452, 522, 422, 348, 386, 482, 552, 450, 360, 544, 344, 378, 452, 522, 422, 348, 386, 482, 552, 450},
				},
				{
					Odds:           []int{55, 75, 65, 3, 7, 4, 25, 45, 35, 9, 15, 11},
					PositionWeight: []int{5789, 2353, 2062, 144607, 60495, 109765, 15711, 7204, 9603, 47340, 27202, 38800, 5789, 2353, 2062, 144607, 60495, 109765, 15711, 7204, 9603, 47340, 27202, 38800},
					SpecialRate:    []int{544, 344, 360, 452, 522, 378, 348, 386, 422, 552, 450, 482, 544, 344, 360, 452, 522, 378, 348, 386, 422, 552, 450, 482},
				},
				{
					Odds:           []int{55, 65, 75, 3, 4, 7, 25, 35, 45, 9, 11, 15},
					PositionWeight: []int{5824, 3807, 819, 144217, 108635, 61615, 15691, 10208, 6644, 47170, 38399, 27702, 5824, 3807, 819, 144217, 108635, 61615, 15691, 10208, 6644, 47170, 38399, 27702},
					SpecialRate:    []int{568, 390, 344, 418, 378, 538, 348, 422, 386, 552, 482, 450, 568, 390, 344, 418, 378, 538, 348, 422, 386, 552, 482, 450},
				},
			},
			BetTable: make([]int64, 12),
		}}
}

// 主游戏逻辑
// 玩法说明：玩家进入游戏，在可下注环节选择任意下注选项进行下注，下注结束后系统开奖，
// 若玩家押中对应的选项，则中奖，系统按照玩家押的下注选项赔率进行派奖；若未押中，则玩家下注的筹码归系统所有。
func (s *CarGodService) MainGame() GameResult {
	eventID := -1

	// 随机levelset
	levelIndex := rand.Intn(len(s.Car.LevelSet))
	levelSet := s.Car.LevelSet[levelIndex]

	// 获取主游戏开奖位置 (0-23)
	// 轮盘一共24个位置，通过权重随机选择一个位置作为本局开奖结果
	mainCarPositionIdx := s.weightGetIdx(levelSet.PositionWeight)

	// 位置对应的车辆图标 id (0-11)
	// 押注选项：一共有四种图标（法拉利、奔驰、宝马和奥迪），其中每个图标有三种颜色（红黄绿），
	// 一共12个下注选项，每个下注选项的赔率不同
	hitCar := s.Car.GroupSet.SymbolPosition[mainCarPositionIdx]

	// 车辆图标对应的基础赔率倍数
	// 每局开始时每个图标的赔率会不同，一共有36种组合，每局开始时随机从36个组合中选择一种供玩家下注
	mainWinOdds := levelSet.Odds[hitCar]

	// 玩家在该车辆图标上的下注金额
	// 如果玩家没有在该图标上下注，则为0
	hitBet := s.Car.BetTable[hitCar]

	// 判断特殊玩法是否触发
	// 根据当前命中车辆的特殊触发率，随机决定是否触发额外的特殊玩法
	eventTrigger := 0
	if rand.Intn(10001) < levelSet.SpecialRate[hitCar] {
		eventTrigger = 1
	}

	// 特殊玩法结果处理
	var eventResult []*EventResult
	if eventTrigger == 1 {
		// 根据权重随机选择一种特殊玩法类型
		eventID = s.weightGetIdx(s.Car.GroupSet.EventWeight)
		// 执行对应的特殊玩法逻辑
		eventResult = s.eventGame(mainCarPositionIdx, eventID, levelSet)
	}

	// 计算主游戏奖金
	// 基础奖金 = 赔率 × 下注金额
	var mainGameWin int64 = 0
	if hitBet != 0 {
		mainGameWin = int64(mainWinOdds) * hitBet
	}

	// 计算特殊玩法奖金和累积奖池奖金
	var eventWin, jackpotWin int64 = 0, 0
	if len(eventResult) > 0 {
		if eventID != EventJackpot {
			// 非累积奖池的特殊玩法：计算额外中奖车辆的奖金
			for _, result := range eventResult {
				bonusHitCar := result.BonusHitCar
				if s.Car.BetTable[bonusHitCar] > 0 {
					eventWin += int64(result.BonusWinOdds) * s.Car.BetTable[bonusHitCar]
				}
			}
		} else {
			// 累积奖池特殊玩法：按照原始下注金额计算奖池奖金
			if hitBet != 0 {
				for _, result := range eventResult {
					jackpotWin += int64(result.BonusWinOdds) * hitBet
				}
			}
		}
	}

	return GameResult{
		LevelIndex:         levelIndex,
		MainCarPositionIdx: mainCarPositionIdx,                  // 主游戏开奖位置
		HitCar:             hitCar,                              // 命中的车辆图标
		MainGameOdds:       mainWinOdds,                         // 主游戏赔率
		EventID:            eventID,                             // 触发的特殊玩法ID
		EventResult:        eventResult,                         // 特殊玩法详细结果
		MainGameWin:        mainGameWin,                         // 主游戏奖金
		EventWin:           eventWin,                            // 特殊玩法奖金
		JackpotWin:         jackpotWin,                          // 累积奖池奖金
		TotalWin:           mainGameWin + eventWin + jackpotWin, // 总奖金
	}
}

// 特殊玩法游戏逻辑 当触发特殊玩法时，根据不同的玩法类型执行相应的奖励逻辑
func (s *CarGodService) eventGame(mainCarPositionIdx, eventID int, levelSet *LevelSet) []*EventResult {
	var eventResult []*EventResult

	switch eventID {
	case EventFortune:
		// 福气奖：额外获得一个最高赔率车辆的奖励
		maxOdds := s.getMaxValue(levelSet.Odds)
		maxOddsIdx := s.getAimIdx(levelSet.Odds, maxOdds)[0]
		maxOddsPositions := s.getAimIdxInt(s.Car.GroupSet.SymbolPosition, maxOddsIdx)

		// 排除主游戏已中奖位置，避免重复计算
		maxOddsPositions = s.removeElement(maxOddsPositions, mainCarPositionIdx)

		if len(maxOddsPositions) > 0 {
			// 随机选择一个最高赔率车辆位置
			bonusHits := maxOddsPositions[rand.Intn(len(maxOddsPositions))]
			bonusHitCar := s.Car.GroupSet.SymbolPosition[bonusHits]
			bonusWinOdds := levelSet.Odds[bonusHitCar]

			eventResult = append(eventResult, &EventResult{
				BonusHits:    bonusHits,
				BonusHitCar:  bonusHitCar,
				BonusWinOdds: bonusWinOdds,
			})
		}
	case EventDouble:
		// 加倍奖：额外获得一个与主游戏相同品牌车辆的奖励
		carTypeHit := s.Car.CarType[mainCarPositionIdx]
		chooseList := s.getAimIdxString(s.Car.CarType, carTypeHit)
		chooseList = s.removeElement(chooseList, mainCarPositionIdx)

		if len(chooseList) > 0 {
			// 随机选择一个相同品牌的车辆位置
			bonusHits := chooseList[rand.Intn(len(chooseList))]
			bonusHitCar := s.Car.GroupSet.SymbolPosition[bonusHits]
			bonusWinOdds := levelSet.Odds[bonusHitCar]

			eventResult = append(eventResult, &EventResult{
				BonusHits:    bonusHits,
				BonusHitCar:  bonusHitCar,
				BonusWinOdds: bonusWinOdds,
			})
		}
	case EventChainLight1, EventChainLight2, EventChainLight3:
		// 连灯奖：按逆时针方向连续点亮1/2/3个相邻位置
		var bonusHits []int

		switch eventID {
		case EventChainLight1:
			// 连灯奖1：点亮逆时针方向的1个相邻位置
			pos := mainCarPositionIdx - 1
			if pos < 0 {
				pos = mainCarPositionIdx + 23 // 轮盘是环形的，0的前一个是23
			}
			bonusHits = append(bonusHits, pos)
		case EventChainLight2:
			// 连灯奖2：点亮逆时针方向的2个相邻位置
			pos1 := mainCarPositionIdx - 1
			if pos1 < 0 {
				pos1 = mainCarPositionIdx + 23
			}
			pos2 := mainCarPositionIdx - 2
			if pos2 < 0 {
				pos2 = mainCarPositionIdx + 22
			}
			bonusHits = append(bonusHits, pos1, pos2)
		case EventChainLight3:
			// 连灯奖3：点亮逆时针方向的3个相邻位置
			pos1 := mainCarPositionIdx - 1
			if pos1 < 0 {
				pos1 = mainCarPositionIdx + 23
			}
			pos2 := mainCarPositionIdx - 2
			if pos2 < 0 {
				pos2 = mainCarPositionIdx + 22
			}
			pos3 := mainCarPositionIdx - 3
			if pos3 < 0 {
				pos3 = mainCarPositionIdx + 21
			}
			bonusHits = append(bonusHits, pos1, pos2, pos3)
		}

		// 计算每个连灯位置的奖励
		for _, bonusHit := range bonusHits {
			bonusHitCar := s.Car.GroupSet.SymbolPosition[bonusHit]
			bonusWinOdds := levelSet.Odds[bonusHitCar]

			eventResult = append(eventResult, &EventResult{
				BonusHits:    bonusHit,
				BonusHitCar:  bonusHitCar,
				BonusWinOdds: bonusWinOdds,
			})
		}
	case EventExtraBonus1, EventExtraBonus2, EventExtraBonus3:
		// 锦上添花：随机选择1/2/3个额外位置获得奖励
		chooseList := make([]int, 24)
		for i := range 24 {
			chooseList[i] = i
		}
		chooseList = s.removeElement(chooseList, mainCarPositionIdx)

		// 根据玩法类型确定随机选择的位置数量
		sampleCount := eventID - 4 // EventExtraBonus1=5, 5-4=1个位置
		bonusHits := s.randomSample(chooseList, sampleCount)

		for _, bonusHit := range bonusHits {
			bonusHitCar := s.Car.GroupSet.SymbolPosition[bonusHit]
			bonusWinOdds := levelSet.Odds[bonusHitCar]

			eventResult = append(eventResult, &EventResult{
				BonusHits:    bonusHit,
				BonusHitCar:  bonusHitCar,
				BonusWinOdds: bonusWinOdds,
			})
		}
	case EventWorldTravel:
		// 云游四海：在轮盘的四个象限中各随机选择一个位置
		bonusHits := []int{
			rand.Intn(6),      // 第一象限：位置0-5
			6 + rand.Intn(6),  // 第二象限：位置6-11
			12 + rand.Intn(6), // 第三象限：位置12-17
			18 + rand.Intn(6), // 第四象限：位置18-23
		}

		bonusHits = s.removeElement(bonusHits, mainCarPositionIdx)

		for _, bonusHit := range bonusHits {
			bonusHitCar := s.Car.GroupSet.SymbolPosition[bonusHit]
			bonusWinOdds := levelSet.Odds[bonusHitCar]

			eventResult = append(eventResult, &EventResult{
				BonusHits:    bonusHit,
				BonusHitCar:  bonusHitCar,
				BonusWinOdds: bonusWinOdds,
			})
		}
	case EventRedCars, EventGreenCars, EventYellowCars:
		// 四大名车：获得特定颜色的四种品牌车辆奖励
		var bonusHits []int
		switch eventID {
		case EventRedCars: // 红色四大名车：位置12,15,18,21
			bonusHits = []int{12, 15, 18, 21}
		case EventGreenCars: // 绿色四大名车：位置13,16,19,22
			bonusHits = []int{13, 16, 19, 22}
		case EventYellowCars: // 黄色四大名车：位置14,17,20,23
			bonusHits = []int{14, 17, 20, 23}
		}

		for _, bonusHit := range bonusHits {
			bonusHitCar := s.Car.GroupSet.SymbolPosition[bonusHit]
			bonusWinOdds := levelSet.Odds[bonusHitCar]

			eventResult = append(eventResult, &EventResult{
				BonusHits:    bonusHit,
				BonusHitCar:  bonusHitCar,
				BonusWinOdds: bonusWinOdds,
			})
		}
	case EventFerrari, EventBenz, EventBMW, EventAudi:
		// 品牌奖：获得特定品牌所有颜色车辆的奖励
		var bonusHits []int
		switch eventID {
		case EventFerrari:
			bonusHits = s.getAimIdxString(s.Car.CarType, "Ferrari")
		case EventBenz:
			bonusHits = s.getAimIdxString(s.Car.CarType, "Benz")
		case EventBMW:
			bonusHits = s.getAimIdxString(s.Car.CarType, "BMW")
		case EventAudi:
			bonusHits = s.getAimIdxString(s.Car.CarType, "Audi")
		}

		bonusHits = s.removeElement(bonusHits, mainCarPositionIdx)

		for _, bonusHit := range bonusHits {
			bonusHitCar := s.Car.GroupSet.SymbolPosition[bonusHit]
			bonusWinOdds := levelSet.Odds[bonusHitCar]

			eventResult = append(eventResult, &EventResult{
				BonusHits:    bonusHit,
				BonusHitCar:  bonusHitCar,
				BonusWinOdds: bonusWinOdds,
			})
		}
	case EventRedLight, EventGreenLight, EventYellowLight:
		// 红灯奖/绿灯奖/黄灯奖：获得特定颜色所有品牌车辆的奖励
		var bonusHits []int
		switch eventID {
		case EventRedLight:
			bonusHits = s.getAimIdxString(s.Car.CarColor, "red")
		case EventGreenLight:
			bonusHits = s.getAimIdxString(s.Car.CarColor, "green")
		case EventYellowLight:
			bonusHits = s.getAimIdxString(s.Car.CarColor, "yellow")
		}

		bonusHits = s.removeElement(bonusHits, mainCarPositionIdx)

		for _, bonusHit := range bonusHits {
			bonusHitCar := s.Car.GroupSet.SymbolPosition[bonusHit]
			bonusWinOdds := levelSet.Odds[bonusHitCar]

			eventResult = append(eventResult, &EventResult{
				BonusHits:    bonusHit,
				BonusHitCar:  bonusHitCar,
				BonusWinOdds: bonusWinOdds,
			})
		}
	case EventGrandSlam:
		// 满贯列车：获得轮盘上除主游戏位置外所有位置的奖励
		bonusHits := make([]int, 24)
		for i := 0; i < 24; i++ {
			bonusHits[i] = i
		}
		bonusHits = s.removeElement(bonusHits, mainCarPositionIdx)

		for _, bonusHit := range bonusHits {
			bonusHitCar := s.Car.GroupSet.SymbolPosition[bonusHit]
			bonusWinOdds := levelSet.Odds[bonusHitCar]

			eventResult = append(eventResult, &EventResult{
				BonusHits:    bonusHit,
				BonusHitCar:  bonusHitCar,
				BonusWinOdds: bonusWinOdds,
			})
		}
	case EventJackpot:
		// 累积奖池：根据权重随机选择奖池档次，获得高倍率奖励
		getJpIdx := s.weightGetIdx(s.Car.GroupSet.JackpotWeight)
		jpOdds := s.Car.GroupSet.JackpotOdds[getJpIdx]

		eventResult = append(eventResult, &EventResult{
			BonusHits:    getJpIdx,
			BonusHitCar:  21, // 固定为21，表示累积奖池
			BonusWinOdds: jpOdds,
		})
	}

	return eventResult
}

// 根据权重随机选择索引 权重越高，被选中的概率越大
func (s *CarGodService) weightGetIdx(weightList []int) int {
	totalWeight := 0
	for _, w := range weightList {
		totalWeight += w
	}

	randNum := rand.Intn(totalWeight)
	sum := 0
	for idx, weight := range weightList {
		sum += weight
		if randNum < sum {
			return idx
		}
	}
	return 0
}

// 获取目标值在整型数组中的所有索引位置
func (s *CarGodService) getAimIdx(symList []int, aimSym int) []int {
	var result []int
	for idx, sym := range symList {
		if sym == aimSym {
			result = append(result, idx)
		}
	}
	return result
}

// 获取目标值在整型数组中的所有索引位置（与getAimIdx功能相同）
func (s *CarGodService) getAimIdxInt(symList []int, aimSym int) []int {
	var result []int
	for idx, sym := range symList {
		if sym == aimSym {
			result = append(result, idx)
		}
	}
	return result
}

// 获取目标值在字符串数组中的所有索引位置
func (s *CarGodService) getAimIdxString(symList []string, aimSym string) []int {
	var result []int
	for idx, sym := range symList {
		if sym == aimSym {
			result = append(result, idx)
		}
	}
	return result
}

// 获取整型数组中的最大值
func (s *CarGodService) getMaxValue(list []int) int {
	max := list[0]
	for _, v := range list {
		if v > max {
			max = v
		}
	}
	return max
}

// 从整型数组中移除指定元素
func (s *CarGodService) removeElement(slice []int, element int) []int {
	var result []int
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}

// 从整型数组中随机采样指定数量的元素
func (s *CarGodService) randomSample(slice []int, count int) []int {
	if count >= len(slice) {
		return slice
	}

	result := make([]int, count)
	indices := rand.Perm(len(slice))
	for i := range count {
		result[i] = slice[indices[i]]
	}
	return result
}
