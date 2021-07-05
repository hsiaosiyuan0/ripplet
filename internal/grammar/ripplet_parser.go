// Code generated from ./grammar/RippletParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // RippletParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 66, 469,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 3, 2, 7, 2, 104, 10, 2, 12, 2, 14, 2, 107, 11,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	3, 120, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 149, 10, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 7, 4, 180, 10, 4, 12, 4, 14, 4, 183, 11, 4, 3, 5, 3, 5, 3, 5,
	5, 5, 188, 10, 5, 3, 5, 5, 5, 191, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	7, 6, 198, 10, 6, 12, 6, 14, 6, 201, 11, 6, 3, 7, 3, 7, 3, 7, 5, 7, 206,
	10, 7, 3, 8, 3, 8, 3, 9, 6, 9, 211, 10, 9, 13, 9, 14, 9, 212, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 16, 7, 16, 240, 10, 16, 12, 16, 14, 16, 243, 11, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 5, 18, 252, 10, 18, 3, 19, 3,
	19, 3, 19, 7, 19, 257, 10, 19, 12, 19, 14, 19, 260, 11, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 5, 21, 269, 10, 21, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 5, 25, 282,
	10, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 7, 26, 289, 10, 26, 12, 26,
	14, 26, 292, 11, 26, 3, 26, 3, 26, 5, 26, 296, 10, 26, 3, 26, 5, 26, 299,
	10, 26, 3, 27, 3, 27, 7, 27, 303, 10, 27, 12, 27, 14, 27, 306, 11, 27,
	3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 7, 29, 315, 10, 29, 12,
	29, 14, 29, 318, 11, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 5, 30, 325,
	10, 30, 3, 30, 3, 30, 3, 30, 5, 30, 330, 10, 30, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 7, 32, 341, 10, 32, 12, 32, 14,
	32, 344, 11, 32, 3, 32, 5, 32, 347, 10, 32, 5, 32, 349, 10, 32, 3, 32,
	3, 32, 3, 33, 5, 33, 354, 10, 33, 3, 33, 3, 33, 5, 33, 358, 10, 33, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 365, 10, 34, 3, 35, 3, 35, 6, 35,
	369, 10, 35, 13, 35, 14, 35, 370, 3, 35, 3, 35, 7, 35, 375, 10, 35, 12,
	35, 14, 35, 378, 11, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38,
	5, 38, 387, 10, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3,
	42, 3, 42, 7, 42, 398, 10, 42, 12, 42, 14, 42, 401, 11, 42, 3, 42, 3, 42,
	3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3,
	46, 3, 46, 7, 46, 417, 10, 46, 12, 46, 14, 46, 420, 11, 46, 7, 46, 422,
	10, 46, 12, 46, 14, 46, 425, 11, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47,
	3, 47, 7, 47, 433, 10, 47, 12, 47, 14, 47, 436, 11, 47, 5, 47, 438, 10,
	47, 3, 47, 5, 47, 441, 10, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48,
	3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 5, 48, 455, 10, 48, 3, 49, 3,
	49, 3, 49, 5, 49, 460, 10, 49, 3, 50, 3, 50, 3, 50, 5, 50, 465, 10, 50,
	3, 51, 3, 51, 3, 51, 2, 3, 6, 52, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
	60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94,
	96, 98, 100, 2, 7, 3, 2, 39, 41, 3, 2, 36, 37, 3, 2, 17, 18, 4, 2, 60,
	60, 62, 62, 5, 2, 3, 9, 11, 11, 13, 15, 2, 491, 2, 105, 3, 2, 2, 2, 4,
	119, 3, 2, 2, 2, 6, 148, 3, 2, 2, 2, 8, 184, 3, 2, 2, 2, 10, 194, 3, 2,
	2, 2, 12, 202, 3, 2, 2, 2, 14, 207, 3, 2, 2, 2, 16, 210, 3, 2, 2, 2, 18,
	214, 3, 2, 2, 2, 20, 218, 3, 2, 2, 2, 22, 222, 3, 2, 2, 2, 24, 226, 3,
	2, 2, 2, 26, 228, 3, 2, 2, 2, 28, 230, 3, 2, 2, 2, 30, 236, 3, 2, 2, 2,
	32, 244, 3, 2, 2, 2, 34, 251, 3, 2, 2, 2, 36, 253, 3, 2, 2, 2, 38, 261,
	3, 2, 2, 2, 40, 268, 3, 2, 2, 2, 42, 270, 3, 2, 2, 2, 44, 275, 3, 2, 2,
	2, 46, 277, 3, 2, 2, 2, 48, 279, 3, 2, 2, 2, 50, 298, 3, 2, 2, 2, 52, 300,
	3, 2, 2, 2, 54, 309, 3, 2, 2, 2, 56, 312, 3, 2, 2, 2, 58, 321, 3, 2, 2,
	2, 60, 331, 3, 2, 2, 2, 62, 336, 3, 2, 2, 2, 64, 353, 3, 2, 2, 2, 66, 364,
	3, 2, 2, 2, 68, 366, 3, 2, 2, 2, 70, 379, 3, 2, 2, 2, 72, 381, 3, 2, 2,
	2, 74, 386, 3, 2, 2, 2, 76, 388, 3, 2, 2, 2, 78, 390, 3, 2, 2, 2, 80, 392,
	3, 2, 2, 2, 82, 394, 3, 2, 2, 2, 84, 404, 3, 2, 2, 2, 86, 406, 3, 2, 2,
	2, 88, 410, 3, 2, 2, 2, 90, 412, 3, 2, 2, 2, 92, 428, 3, 2, 2, 2, 94, 454,
	3, 2, 2, 2, 96, 459, 3, 2, 2, 2, 98, 464, 3, 2, 2, 2, 100, 466, 3, 2, 2,
	2, 102, 104, 5, 4, 3, 2, 103, 102, 3, 2, 2, 2, 104, 107, 3, 2, 2, 2, 105,
	103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 108, 3, 2, 2, 2, 107, 105,
	3, 2, 2, 2, 108, 109, 7, 2, 2, 3, 109, 3, 3, 2, 2, 2, 110, 120, 5, 26,
	14, 2, 111, 120, 5, 58, 30, 2, 112, 120, 5, 60, 31, 2, 113, 120, 5, 28,
	15, 2, 114, 120, 5, 20, 11, 2, 115, 120, 5, 22, 12, 2, 116, 120, 5, 8,
	5, 2, 117, 120, 5, 42, 22, 2, 118, 120, 5, 56, 29, 2, 119, 110, 3, 2, 2,
	2, 119, 111, 3, 2, 2, 2, 119, 112, 3, 2, 2, 2, 119, 113, 3, 2, 2, 2, 119,
	114, 3, 2, 2, 2, 119, 115, 3, 2, 2, 2, 119, 116, 3, 2, 2, 2, 119, 117,
	3, 2, 2, 2, 119, 118, 3, 2, 2, 2, 120, 5, 3, 2, 2, 2, 121, 122, 8, 4, 1,
	2, 122, 123, 7, 10, 2, 2, 123, 149, 5, 6, 4, 22, 124, 125, 7, 12, 2, 2,
	125, 149, 5, 6, 4, 19, 126, 127, 5, 46, 24, 2, 127, 128, 5, 62, 32, 2,
	128, 149, 3, 2, 2, 2, 129, 130, 5, 48, 25, 2, 130, 131, 7, 50, 2, 2, 131,
	132, 5, 52, 27, 2, 132, 149, 3, 2, 2, 2, 133, 134, 7, 17, 2, 2, 134, 149,
	5, 6, 4, 11, 135, 136, 7, 18, 2, 2, 136, 149, 5, 6, 4, 10, 137, 149, 7,
	21, 2, 2, 138, 149, 5, 46, 24, 2, 139, 149, 5, 90, 46, 2, 140, 149, 5,
	92, 47, 2, 141, 149, 5, 66, 34, 2, 142, 143, 7, 30, 2, 2, 143, 144, 5,
	6, 4, 2, 144, 145, 7, 31, 2, 2, 145, 149, 3, 2, 2, 2, 146, 147, 7, 30,
	2, 2, 147, 149, 7, 31, 2, 2, 148, 121, 3, 2, 2, 2, 148, 124, 3, 2, 2, 2,
	148, 126, 3, 2, 2, 2, 148, 129, 3, 2, 2, 2, 148, 133, 3, 2, 2, 2, 148,
	135, 3, 2, 2, 2, 148, 137, 3, 2, 2, 2, 148, 138, 3, 2, 2, 2, 148, 139,
	3, 2, 2, 2, 148, 140, 3, 2, 2, 2, 148, 141, 3, 2, 2, 2, 148, 142, 3, 2,
	2, 2, 148, 146, 3, 2, 2, 2, 149, 181, 3, 2, 2, 2, 150, 151, 12, 21, 2,
	2, 151, 152, 7, 11, 2, 2, 152, 180, 5, 6, 4, 22, 153, 154, 12, 20, 2, 2,
	154, 155, 7, 13, 2, 2, 155, 180, 5, 6, 4, 21, 156, 157, 12, 16, 2, 2, 157,
	158, 7, 38, 2, 2, 158, 180, 5, 6, 4, 16, 159, 160, 12, 15, 2, 2, 160, 161,
	9, 2, 2, 2, 161, 180, 5, 6, 4, 16, 162, 163, 12, 14, 2, 2, 163, 164, 9,
	3, 2, 2, 164, 180, 5, 6, 4, 15, 165, 166, 12, 13, 2, 2, 166, 167, 7, 14,
	2, 2, 167, 180, 5, 6, 4, 14, 168, 169, 12, 12, 2, 2, 169, 170, 7, 15, 2,
	2, 170, 180, 5, 6, 4, 13, 171, 172, 12, 24, 2, 2, 172, 173, 7, 28, 2, 2,
	173, 174, 5, 6, 4, 2, 174, 175, 7, 29, 2, 2, 175, 180, 3, 2, 2, 2, 176,
	177, 12, 23, 2, 2, 177, 178, 7, 35, 2, 2, 178, 180, 7, 58, 2, 2, 179, 150,
	3, 2, 2, 2, 179, 153, 3, 2, 2, 2, 179, 156, 3, 2, 2, 2, 179, 159, 3, 2,
	2, 2, 179, 162, 3, 2, 2, 2, 179, 165, 3, 2, 2, 2, 179, 168, 3, 2, 2, 2,
	179, 171, 3, 2, 2, 2, 179, 176, 3, 2, 2, 2, 180, 183, 3, 2, 2, 2, 181,
	179, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 7, 3, 2, 2, 2, 183, 181, 3,
	2, 2, 2, 184, 185, 7, 20, 2, 2, 185, 187, 7, 26, 2, 2, 186, 188, 5, 10,
	6, 2, 187, 186, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 190, 3, 2, 2, 2,
	189, 191, 5, 16, 9, 2, 190, 189, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191,
	192, 3, 2, 2, 2, 192, 193, 7, 27, 2, 2, 193, 9, 3, 2, 2, 2, 194, 199, 5,
	12, 7, 2, 195, 196, 7, 32, 2, 2, 196, 198, 5, 12, 7, 2, 197, 195, 3, 2,
	2, 2, 198, 201, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2,
	200, 11, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 202, 205, 7, 58, 2, 2, 203,
	204, 7, 43, 2, 2, 204, 206, 5, 14, 8, 2, 205, 203, 3, 2, 2, 2, 205, 206,
	3, 2, 2, 2, 206, 13, 3, 2, 2, 2, 207, 208, 5, 6, 4, 2, 208, 15, 3, 2, 2,
	2, 209, 211, 5, 18, 10, 2, 210, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2,
	212, 210, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 17, 3, 2, 2, 2, 214, 215,
	7, 58, 2, 2, 215, 216, 5, 48, 25, 2, 216, 217, 5, 56, 29, 2, 217, 19, 3,
	2, 2, 2, 218, 219, 5, 46, 24, 2, 219, 220, 7, 43, 2, 2, 220, 221, 5, 4,
	3, 2, 221, 21, 3, 2, 2, 2, 222, 223, 5, 24, 13, 2, 223, 224, 7, 42, 2,
	2, 224, 225, 5, 4, 3, 2, 225, 23, 3, 2, 2, 2, 226, 227, 5, 46, 24, 2, 227,
	25, 3, 2, 2, 2, 228, 229, 5, 6, 4, 2, 229, 27, 3, 2, 2, 2, 230, 231, 7,
	6, 2, 2, 231, 232, 5, 6, 4, 2, 232, 233, 7, 26, 2, 2, 233, 234, 5, 30,
	16, 2, 234, 235, 7, 27, 2, 2, 235, 29, 3, 2, 2, 2, 236, 241, 5, 32, 17,
	2, 237, 238, 7, 32, 2, 2, 238, 240, 5, 32, 17, 2, 239, 237, 3, 2, 2, 2,
	240, 243, 3, 2, 2, 2, 241, 239, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242,
	31, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 244, 245, 5, 34, 18, 2, 245, 246,
	7, 50, 2, 2, 246, 247, 5, 4, 3, 2, 247, 33, 3, 2, 2, 2, 248, 252, 5, 36,
	19, 2, 249, 252, 5, 38, 20, 2, 250, 252, 7, 52, 2, 2, 251, 248, 3, 2, 2,
	2, 251, 249, 3, 2, 2, 2, 251, 250, 3, 2, 2, 2, 252, 35, 3, 2, 2, 2, 253,
	258, 5, 40, 21, 2, 254, 255, 7, 15, 2, 2, 255, 257, 5, 40, 21, 2, 256,
	254, 3, 2, 2, 2, 257, 260, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2, 258, 259,
	3, 2, 2, 2, 259, 37, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 261, 262, 9, 4,
	2, 2, 262, 263, 7, 30, 2, 2, 263, 264, 5, 46, 24, 2, 264, 265, 7, 31, 2,
	2, 265, 39, 3, 2, 2, 2, 266, 269, 5, 46, 24, 2, 267, 269, 5, 74, 38, 2,
	268, 266, 3, 2, 2, 2, 268, 267, 3, 2, 2, 2, 269, 41, 3, 2, 2, 2, 270, 271,
	7, 7, 2, 2, 271, 272, 5, 44, 23, 2, 272, 273, 5, 48, 25, 2, 273, 274, 5,
	52, 27, 2, 274, 43, 3, 2, 2, 2, 275, 276, 5, 46, 24, 2, 276, 45, 3, 2,
	2, 2, 277, 278, 7, 58, 2, 2, 278, 47, 3, 2, 2, 2, 279, 281, 7, 30, 2, 2,
	280, 282, 5, 50, 26, 2, 281, 280, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282,
	283, 3, 2, 2, 2, 283, 284, 7, 31, 2, 2, 284, 49, 3, 2, 2, 2, 285, 290,
	5, 46, 24, 2, 286, 287, 7, 32, 2, 2, 287, 289, 5, 46, 24, 2, 288, 286,
	3, 2, 2, 2, 289, 292, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 290, 291, 3, 2,
	2, 2, 291, 295, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 293, 294, 7, 32, 2, 2,
	294, 296, 5, 54, 28, 2, 295, 293, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296,
	299, 3, 2, 2, 2, 297, 299, 5, 54, 28, 2, 298, 285, 3, 2, 2, 2, 298, 297,
	3, 2, 2, 2, 299, 51, 3, 2, 2, 2, 300, 304, 7, 26, 2, 2, 301, 303, 5, 4,
	3, 2, 302, 301, 3, 2, 2, 2, 303, 306, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2,
	304, 305, 3, 2, 2, 2, 305, 307, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 307,
	308, 7, 27, 2, 2, 308, 53, 3, 2, 2, 2, 309, 310, 7, 51, 2, 2, 310, 311,
	5, 46, 24, 2, 311, 55, 3, 2, 2, 2, 312, 316, 7, 26, 2, 2, 313, 315, 5,
	4, 3, 2, 314, 313, 3, 2, 2, 2, 315, 318, 3, 2, 2, 2, 316, 314, 3, 2, 2,
	2, 316, 317, 3, 2, 2, 2, 317, 319, 3, 2, 2, 2, 318, 316, 3, 2, 2, 2, 319,
	320, 7, 27, 2, 2, 320, 57, 3, 2, 2, 2, 321, 322, 7, 3, 2, 2, 322, 324,
	5, 6, 4, 2, 323, 325, 7, 4, 2, 2, 324, 323, 3, 2, 2, 2, 324, 325, 3, 2,
	2, 2, 325, 326, 3, 2, 2, 2, 326, 329, 5, 4, 3, 2, 327, 328, 7, 5, 2, 2,
	328, 330, 5, 4, 3, 2, 329, 327, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330,
	59, 3, 2, 2, 2, 331, 332, 7, 8, 2, 2, 332, 333, 5, 4, 3, 2, 333, 334, 7,
	9, 2, 2, 334, 335, 5, 6, 4, 2, 335, 61, 3, 2, 2, 2, 336, 348, 7, 30, 2,
	2, 337, 342, 5, 64, 33, 2, 338, 339, 7, 32, 2, 2, 339, 341, 5, 64, 33,
	2, 340, 338, 3, 2, 2, 2, 341, 344, 3, 2, 2, 2, 342, 340, 3, 2, 2, 2, 342,
	343, 3, 2, 2, 2, 343, 346, 3, 2, 2, 2, 344, 342, 3, 2, 2, 2, 345, 347,
	7, 32, 2, 2, 346, 345, 3, 2, 2, 2, 346, 347, 3, 2, 2, 2, 347, 349, 3, 2,
	2, 2, 348, 337, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 350, 3, 2, 2, 2,
	350, 351, 7, 31, 2, 2, 351, 63, 3, 2, 2, 2, 352, 354, 7, 51, 2, 2, 353,
	352, 3, 2, 2, 2, 353, 354, 3, 2, 2, 2, 354, 357, 3, 2, 2, 2, 355, 358,
	5, 6, 4, 2, 356, 358, 7, 58, 2, 2, 357, 355, 3, 2, 2, 2, 357, 356, 3, 2,
	2, 2, 358, 65, 3, 2, 2, 2, 359, 365, 5, 72, 37, 2, 360, 365, 5, 70, 36,
	2, 361, 365, 5, 74, 38, 2, 362, 365, 5, 68, 35, 2, 363, 365, 5, 82, 42,
	2, 364, 359, 3, 2, 2, 2, 364, 360, 3, 2, 2, 2, 364, 361, 3, 2, 2, 2, 364,
	362, 3, 2, 2, 2, 364, 363, 3, 2, 2, 2, 365, 67, 3, 2, 2, 2, 366, 368, 7,
	25, 2, 2, 367, 369, 9, 5, 2, 2, 368, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2,
	2, 370, 368, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372,
	376, 7, 63, 2, 2, 373, 375, 7, 58, 2, 2, 374, 373, 3, 2, 2, 2, 375, 378,
	3, 2, 2, 2, 376, 374, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 69, 3, 2,
	2, 2, 378, 376, 3, 2, 2, 2, 379, 380, 7, 54, 2, 2, 380, 71, 3, 2, 2, 2,
	381, 382, 7, 53, 2, 2, 382, 73, 3, 2, 2, 2, 383, 387, 5, 76, 39, 2, 384,
	387, 5, 78, 40, 2, 385, 387, 5, 80, 41, 2, 386, 383, 3, 2, 2, 2, 386, 384,
	3, 2, 2, 2, 386, 385, 3, 2, 2, 2, 387, 75, 3, 2, 2, 2, 388, 389, 7, 55,
	2, 2, 389, 77, 3, 2, 2, 2, 390, 391, 7, 56, 2, 2, 391, 79, 3, 2, 2, 2,
	392, 393, 7, 57, 2, 2, 393, 81, 3, 2, 2, 2, 394, 399, 7, 59, 2, 2, 395,
	398, 7, 66, 2, 2, 396, 398, 5, 86, 44, 2, 397, 395, 3, 2, 2, 2, 397, 396,
	3, 2, 2, 2, 398, 401, 3, 2, 2, 2, 399, 397, 3, 2, 2, 2, 399, 400, 3, 2,
	2, 2, 400, 402, 3, 2, 2, 2, 401, 399, 3, 2, 2, 2, 402, 403, 7, 65, 2, 2,
	403, 83, 3, 2, 2, 2, 404, 405, 7, 66, 2, 2, 405, 85, 3, 2, 2, 2, 406, 407,
	7, 64, 2, 2, 407, 408, 5, 88, 45, 2, 408, 409, 7, 27, 2, 2, 409, 87, 3,
	2, 2, 2, 410, 411, 5, 6, 4, 2, 411, 89, 3, 2, 2, 2, 412, 423, 7, 28, 2,
	2, 413, 418, 5, 6, 4, 2, 414, 415, 7, 32, 2, 2, 415, 417, 5, 6, 4, 2, 416,
	414, 3, 2, 2, 2, 417, 420, 3, 2, 2, 2, 418, 416, 3, 2, 2, 2, 418, 419,
	3, 2, 2, 2, 419, 422, 3, 2, 2, 2, 420, 418, 3, 2, 2, 2, 421, 413, 3, 2,
	2, 2, 422, 425, 3, 2, 2, 2, 423, 421, 3, 2, 2, 2, 423, 424, 3, 2, 2, 2,
	424, 426, 3, 2, 2, 2, 425, 423, 3, 2, 2, 2, 426, 427, 7, 29, 2, 2, 427,
	91, 3, 2, 2, 2, 428, 437, 7, 26, 2, 2, 429, 434, 5, 94, 48, 2, 430, 431,
	7, 32, 2, 2, 431, 433, 5, 94, 48, 2, 432, 430, 3, 2, 2, 2, 433, 436, 3,
	2, 2, 2, 434, 432, 3, 2, 2, 2, 434, 435, 3, 2, 2, 2, 435, 438, 3, 2, 2,
	2, 436, 434, 3, 2, 2, 2, 437, 429, 3, 2, 2, 2, 437, 438, 3, 2, 2, 2, 438,
	440, 3, 2, 2, 2, 439, 441, 7, 32, 2, 2, 440, 439, 3, 2, 2, 2, 440, 441,
	3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442, 443, 7, 27, 2, 2, 443, 93, 3, 2,
	2, 2, 444, 445, 5, 96, 49, 2, 445, 446, 7, 33, 2, 2, 446, 447, 5, 6, 4,
	2, 447, 455, 3, 2, 2, 2, 448, 449, 7, 28, 2, 2, 449, 450, 5, 6, 4, 2, 450,
	451, 7, 29, 2, 2, 451, 452, 7, 33, 2, 2, 452, 453, 5, 6, 4, 2, 453, 455,
	3, 2, 2, 2, 454, 444, 3, 2, 2, 2, 454, 448, 3, 2, 2, 2, 455, 95, 3, 2,
	2, 2, 456, 460, 7, 58, 2, 2, 457, 460, 5, 82, 42, 2, 458, 460, 5, 74, 38,
	2, 459, 456, 3, 2, 2, 2, 459, 457, 3, 2, 2, 2, 459, 458, 3, 2, 2, 2, 460,
	97, 3, 2, 2, 2, 461, 465, 5, 100, 51, 2, 462, 465, 5, 72, 37, 2, 463, 465,
	5, 70, 36, 2, 464, 461, 3, 2, 2, 2, 464, 462, 3, 2, 2, 2, 464, 463, 3,
	2, 2, 2, 465, 99, 3, 2, 2, 2, 466, 467, 9, 6, 2, 2, 467, 101, 3, 2, 2,
	2, 43, 105, 119, 148, 179, 181, 187, 190, 199, 205, 212, 241, 251, 258,
	268, 281, 290, 295, 298, 304, 316, 324, 329, 342, 346, 348, 353, 357, 364,
	370, 376, 386, 397, 399, 418, 423, 434, 437, 440, 454, 459, 464,
}
var literalNames = []string{
	"", "'if'", "'then'", "'else'", "'match'", "'fn'", "'repeat'", "'until'",
	"'typeof'", "'is'", "'not'", "'isnt'", "'and'", "'or'", "'break'", "'ok'",
	"'err'", "'return'", "'object'", "'this'", "", "", "", "", "'{'", "'}'",
	"'['", "']'", "'('", "')'", "','", "':'", "';'", "'.'", "'+'", "'-'", "'**'",
	"'*'", "", "'%'", "':='", "'='", "'<'", "'>'", "'=='", "'!='", "'<='",
	"'>='", "'=>'", "'...'", "'_'", "'nil'",
}
var symbolicNames = []string{
	"", "If", "Then", "Else", "Match", "Fn", "Repeat", "Until", "Typeof", "Is",
	"Not", "IsNot", "And", "Or", "Break", "Ok", "Err", "Return", "Object",
	"This", "LineTerminator", "Whitespace", "Comment", "RegexpStart", "BraceOpen",
	"BraceClose", "BracketOpen", "BracketClose", "ParenOpen", "ParenClose",
	"Comma", "Colon", "SemiColon", "Dot", "Plus", "Minus", "Power", "Multiply",
	"Divide", "Modulus", "Declare", "Assign", "LessThan", "MoreThan", "Equals",
	"NotEquals", "LessThanEquals", "GreaterThanEquals", "LambdaConnect", "Ellipsis",
	"Discard", "NilLiteral", "BooLiteral", "IntLiteral", "HexLiteral", "RealLiteral",
	"Identifier", "StringOpen", "RegexpComment", "RegexpNewline", "RegexpContent",
	"RexexpClose", "StringInterpStart", "StringClose", "StringQuoted",
}

var ruleNames = []string{
	"program", "statement", "expression", "objDeclareStmt", "objProps", "objProp",
	"objPropInit", "objMethods", "objMethod", "assignStmt", "varDeclareStmt",
	"varDeclareLhs", "exprStmt", "matchStmt", "mathClauses", "mathClause",
	"matchClauseTest", "matchClauseTestVal", "matchClauseTestUnwrap", "matchClauseVal",
	"fnDeclareStmt", "fnName", "identifer", "formalParams", "formalParamList",
	"fnBody", "restParamArg", "blockStmt", "ifStmt", "repeatStmt", "arguments",
	"argument", "literal", "regexpLiteral", "boolLiteral", "nilLiteral", "numberLiteral",
	"intLiteral", "hexLiteral", "realLiteral", "stringLiteral", "stringQuoted",
	"stringInterp", "stringInterpExpr", "arrayLiteral", "objectLiteral", "prop",
	"propName", "reservedWord", "keyword",
}

type RippletParser struct {
	*antlr.BaseParser
}

// NewRippletParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *RippletParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewRippletParser(input antlr.TokenStream) *RippletParser {
	this := new(RippletParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "RippletParser.g4"

	return this
}

// RippletParser tokens.
const (
	RippletParserEOF               = antlr.TokenEOF
	RippletParserIf                = 1
	RippletParserThen              = 2
	RippletParserElse              = 3
	RippletParserMatch             = 4
	RippletParserFn                = 5
	RippletParserRepeat            = 6
	RippletParserUntil             = 7
	RippletParserTypeof            = 8
	RippletParserIs                = 9
	RippletParserNot               = 10
	RippletParserIsNot             = 11
	RippletParserAnd               = 12
	RippletParserOr                = 13
	RippletParserBreak             = 14
	RippletParserOk                = 15
	RippletParserErr               = 16
	RippletParserReturn            = 17
	RippletParserObject            = 18
	RippletParserThis              = 19
	RippletParserLineTerminator    = 20
	RippletParserWhitespace        = 21
	RippletParserComment           = 22
	RippletParserRegexpStart       = 23
	RippletParserBraceOpen         = 24
	RippletParserBraceClose        = 25
	RippletParserBracketOpen       = 26
	RippletParserBracketClose      = 27
	RippletParserParenOpen         = 28
	RippletParserParenClose        = 29
	RippletParserComma             = 30
	RippletParserColon             = 31
	RippletParserSemiColon         = 32
	RippletParserDot               = 33
	RippletParserPlus              = 34
	RippletParserMinus             = 35
	RippletParserPower             = 36
	RippletParserMultiply          = 37
	RippletParserDivide            = 38
	RippletParserModulus           = 39
	RippletParserDeclare           = 40
	RippletParserAssign            = 41
	RippletParserLessThan          = 42
	RippletParserMoreThan          = 43
	RippletParserEquals            = 44
	RippletParserNotEquals         = 45
	RippletParserLessThanEquals    = 46
	RippletParserGreaterThanEquals = 47
	RippletParserLambdaConnect     = 48
	RippletParserEllipsis          = 49
	RippletParserDiscard           = 50
	RippletParserNilLiteral        = 51
	RippletParserBooLiteral        = 52
	RippletParserIntLiteral        = 53
	RippletParserHexLiteral        = 54
	RippletParserRealLiteral       = 55
	RippletParserIdentifier        = 56
	RippletParserStringOpen        = 57
	RippletParserRegexpComment     = 58
	RippletParserRegexpNewline     = 59
	RippletParserRegexpContent     = 60
	RippletParserRexexpClose       = 61
	RippletParserStringInterpStart = 62
	RippletParserStringClose       = 63
	RippletParserStringQuoted      = 64
)

// RippletParser rules.
const (
	RippletParserRULE_program               = 0
	RippletParserRULE_statement             = 1
	RippletParserRULE_expression            = 2
	RippletParserRULE_objDeclareStmt        = 3
	RippletParserRULE_objProps              = 4
	RippletParserRULE_objProp               = 5
	RippletParserRULE_objPropInit           = 6
	RippletParserRULE_objMethods            = 7
	RippletParserRULE_objMethod             = 8
	RippletParserRULE_assignStmt            = 9
	RippletParserRULE_varDeclareStmt        = 10
	RippletParserRULE_varDeclareLhs         = 11
	RippletParserRULE_exprStmt              = 12
	RippletParserRULE_matchStmt             = 13
	RippletParserRULE_mathClauses           = 14
	RippletParserRULE_mathClause            = 15
	RippletParserRULE_matchClauseTest       = 16
	RippletParserRULE_matchClauseTestVal    = 17
	RippletParserRULE_matchClauseTestUnwrap = 18
	RippletParserRULE_matchClauseVal        = 19
	RippletParserRULE_fnDeclareStmt         = 20
	RippletParserRULE_fnName                = 21
	RippletParserRULE_identifer             = 22
	RippletParserRULE_formalParams          = 23
	RippletParserRULE_formalParamList       = 24
	RippletParserRULE_fnBody                = 25
	RippletParserRULE_restParamArg          = 26
	RippletParserRULE_blockStmt             = 27
	RippletParserRULE_ifStmt                = 28
	RippletParserRULE_repeatStmt            = 29
	RippletParserRULE_arguments             = 30
	RippletParserRULE_argument              = 31
	RippletParserRULE_literal               = 32
	RippletParserRULE_regexpLiteral         = 33
	RippletParserRULE_boolLiteral           = 34
	RippletParserRULE_nilLiteral            = 35
	RippletParserRULE_numberLiteral         = 36
	RippletParserRULE_intLiteral            = 37
	RippletParserRULE_hexLiteral            = 38
	RippletParserRULE_realLiteral           = 39
	RippletParserRULE_stringLiteral         = 40
	RippletParserRULE_stringQuoted          = 41
	RippletParserRULE_stringInterp          = 42
	RippletParserRULE_stringInterpExpr      = 43
	RippletParserRULE_arrayLiteral          = 44
	RippletParserRULE_objectLiteral         = 45
	RippletParserRULE_prop                  = 46
	RippletParserRULE_propName              = 47
	RippletParserRULE_reservedWord          = 48
	RippletParserRULE_keyword               = 49
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(RippletParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *RippletParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RippletParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RippletParserIf)|(1<<RippletParserMatch)|(1<<RippletParserFn)|(1<<RippletParserRepeat)|(1<<RippletParserTypeof)|(1<<RippletParserNot)|(1<<RippletParserOk)|(1<<RippletParserErr)|(1<<RippletParserObject)|(1<<RippletParserThis)|(1<<RippletParserRegexpStart)|(1<<RippletParserBraceOpen)|(1<<RippletParserBracketOpen)|(1<<RippletParserParenOpen))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(RippletParserNilLiteral-51))|(1<<(RippletParserBooLiteral-51))|(1<<(RippletParserIntLiteral-51))|(1<<(RippletParserHexLiteral-51))|(1<<(RippletParserRealLiteral-51))|(1<<(RippletParserIdentifier-51))|(1<<(RippletParserStringOpen-51)))) != 0) {
		{
			p.SetState(100)
			p.Statement()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(106)
		p.Match(RippletParserEOF)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ExprStmt() IExprStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprStmtContext)
}

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) RepeatStmt() IRepeatStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRepeatStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRepeatStmtContext)
}

func (s *StatementContext) MatchStmt() IMatchStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatchStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatchStmtContext)
}

func (s *StatementContext) AssignStmt() IAssignStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *StatementContext) VarDeclareStmt() IVarDeclareStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclareStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclareStmtContext)
}

func (s *StatementContext) ObjDeclareStmt() IObjDeclareStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjDeclareStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjDeclareStmtContext)
}

func (s *StatementContext) FnDeclareStmt() IFnDeclareStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFnDeclareStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFnDeclareStmtContext)
}

func (s *StatementContext) BlockStmt() IBlockStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockStmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *RippletParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RippletParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.ExprStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.IfStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.RepeatStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(111)
			p.MatchStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(112)
			p.AssignStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(113)
			p.VarDeclareStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(114)
			p.ObjDeclareStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(115)
			p.FnDeclareStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(116)
			p.BlockStmt()
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulExprContext struct {
	*ExpressionContext
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MulExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulExprContext) Multiply() antlr.TerminalNode {
	return s.GetToken(RippletParserMultiply, 0)
}

func (s *MulExprContext) Divide() antlr.TerminalNode {
	return s.GetToken(RippletParserDivide, 0)
}

func (s *MulExprContext) Modulus() antlr.TerminalNode {
	return s.GetToken(RippletParserModulus, 0)
}

func (s *MulExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterMulExpr(s)
	}
}

func (s *MulExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitMulExpr(s)
	}
}

type OkExprContext struct {
	*ExpressionContext
}

func NewOkExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OkExprContext {
	var p = new(OkExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OkExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OkExprContext) Ok() antlr.TerminalNode {
	return s.GetToken(RippletParserOk, 0)
}

func (s *OkExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OkExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterOkExpr(s)
	}
}

func (s *OkExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitOkExpr(s)
	}
}

type ErrExprContext struct {
	*ExpressionContext
}

func NewErrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ErrExprContext {
	var p = new(ErrExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ErrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrExprContext) Err() antlr.TerminalNode {
	return s.GetToken(RippletParserErr, 0)
}

func (s *ErrExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ErrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterErrExpr(s)
	}
}

func (s *ErrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitErrExpr(s)
	}
}

type AddExprContext struct {
	*ExpressionContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddExprContext) Plus() antlr.TerminalNode {
	return s.GetToken(RippletParserPlus, 0)
}

func (s *AddExprContext) Minus() antlr.TerminalNode {
	return s.GetToken(RippletParserMinus, 0)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

type LogicAndExprContext struct {
	*ExpressionContext
}

func NewLogicAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicAndExprContext {
	var p = new(LogicAndExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogicAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicAndExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LogicAndExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicAndExprContext) And() antlr.TerminalNode {
	return s.GetToken(RippletParserAnd, 0)
}

func (s *LogicAndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterLogicAndExpr(s)
	}
}

func (s *LogicAndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitLogicAndExpr(s)
	}
}

type VoidExprContext struct {
	*ExpressionContext
}

func NewVoidExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VoidExprContext {
	var p = new(VoidExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *VoidExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoidExprContext) ParenOpen() antlr.TerminalNode {
	return s.GetToken(RippletParserParenOpen, 0)
}

func (s *VoidExprContext) ParenClose() antlr.TerminalNode {
	return s.GetToken(RippletParserParenClose, 0)
}

func (s *VoidExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterVoidExpr(s)
	}
}

func (s *VoidExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitVoidExpr(s)
	}
}

type PowerExprContext struct {
	*ExpressionContext
}

func NewPowerExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerExprContext {
	var p = new(PowerExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PowerExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *PowerExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PowerExprContext) Power() antlr.TerminalNode {
	return s.GetToken(RippletParserPower, 0)
}

func (s *PowerExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterPowerExpr(s)
	}
}

func (s *PowerExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitPowerExpr(s)
	}
}

type ArrayExprContext struct {
	*ExpressionContext
}

func NewArrayExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayExprContext {
	var p = new(ArrayExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ArrayExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExprContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ArrayExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterArrayExpr(s)
	}
}

func (s *ArrayExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitArrayExpr(s)
	}
}

type EqualityIsNotExprContext struct {
	*ExpressionContext
}

func NewEqualityIsNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityIsNotExprContext {
	var p = new(EqualityIsNotExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EqualityIsNotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityIsNotExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqualityIsNotExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqualityIsNotExprContext) IsNot() antlr.TerminalNode {
	return s.GetToken(RippletParserIsNot, 0)
}

func (s *EqualityIsNotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterEqualityIsNotExpr(s)
	}
}

func (s *EqualityIsNotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitEqualityIsNotExpr(s)
	}
}

type MemterIdxExprContext struct {
	*ExpressionContext
}

func NewMemterIdxExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemterIdxExprContext {
	var p = new(MemterIdxExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MemterIdxExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemterIdxExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MemterIdxExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MemterIdxExprContext) BracketOpen() antlr.TerminalNode {
	return s.GetToken(RippletParserBracketOpen, 0)
}

func (s *MemterIdxExprContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(RippletParserBracketClose, 0)
}

func (s *MemterIdxExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterMemterIdxExpr(s)
	}
}

func (s *MemterIdxExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitMemterIdxExpr(s)
	}
}

type MemberDotExprContext struct {
	*ExpressionContext
}

func NewMemberDotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberDotExprContext {
	var p = new(MemberDotExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MemberDotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberDotExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MemberDotExprContext) Dot() antlr.TerminalNode {
	return s.GetToken(RippletParserDot, 0)
}

func (s *MemberDotExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RippletParserIdentifier, 0)
}

func (s *MemberDotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterMemberDotExpr(s)
	}
}

func (s *MemberDotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitMemberDotExpr(s)
	}
}

type TypeofExprContext struct {
	*ExpressionContext
}

func NewTypeofExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeofExprContext {
	var p = new(TypeofExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TypeofExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeofExprContext) Typeof() antlr.TerminalNode {
	return s.GetToken(RippletParserTypeof, 0)
}

func (s *TypeofExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TypeofExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterTypeofExpr(s)
	}
}

func (s *TypeofExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitTypeofExpr(s)
	}
}

type ObjectExprContext struct {
	*ExpressionContext
}

func NewObjectExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectExprContext {
	var p = new(ObjectExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ObjectExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectExprContext) ObjectLiteral() IObjectLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectLiteralContext)
}

func (s *ObjectExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterObjectExpr(s)
	}
}

func (s *ObjectExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitObjectExpr(s)
	}
}

type IdentifierExprContext struct {
	*ExpressionContext
}

func NewIdentifierExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExprContext {
	var p = new(IdentifierExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExprContext) Identifer() IIdentiferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentiferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentiferContext)
}

func (s *IdentifierExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterIdentifierExpr(s)
	}
}

func (s *IdentifierExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitIdentifierExpr(s)
	}
}

type LogicOrExprContext struct {
	*ExpressionContext
}

func NewLogicOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicOrExprContext {
	var p = new(LogicOrExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogicOrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicOrExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LogicOrExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicOrExprContext) Or() antlr.TerminalNode {
	return s.GetToken(RippletParserOr, 0)
}

func (s *LogicOrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterLogicOrExpr(s)
	}
}

func (s *LogicOrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitLogicOrExpr(s)
	}
}

type LiteralExprContext struct {
	*ExpressionContext
}

func NewLiteralExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExprContext {
	var p = new(LiteralExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExprContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterLiteralExpr(s)
	}
}

func (s *LiteralExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitLiteralExpr(s)
	}
}

type CallExprContext struct {
	*ExpressionContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) Identifer() IIdentiferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentiferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentiferContext)
}

func (s *CallExprContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

type NotExprContext struct {
	*ExpressionContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Not() antlr.TerminalNode {
	return s.GetToken(RippletParserNot, 0)
}

func (s *NotExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

type ParenExprContext struct {
	*ExpressionContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) ParenOpen() antlr.TerminalNode {
	return s.GetToken(RippletParserParenOpen, 0)
}

func (s *ParenExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExprContext) ParenClose() antlr.TerminalNode {
	return s.GetToken(RippletParserParenClose, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

type EqualityIsExprContext struct {
	*ExpressionContext
}

func NewEqualityIsExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityIsExprContext {
	var p = new(EqualityIsExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EqualityIsExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityIsExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqualityIsExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqualityIsExprContext) Is() antlr.TerminalNode {
	return s.GetToken(RippletParserIs, 0)
}

func (s *EqualityIsExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterEqualityIsExpr(s)
	}
}

func (s *EqualityIsExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitEqualityIsExpr(s)
	}
}

type ThisExprContext struct {
	*ExpressionContext
}

func NewThisExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThisExprContext {
	var p = new(ThisExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ThisExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThisExprContext) This() antlr.TerminalNode {
	return s.GetToken(RippletParserThis, 0)
}

func (s *ThisExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterThisExpr(s)
	}
}

func (s *ThisExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitThisExpr(s)
	}
}

type FnExprContext struct {
	*ExpressionContext
}

func NewFnExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FnExprContext {
	var p = new(FnExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FnExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnExprContext) FormalParams() IFormalParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParamsContext)
}

func (s *FnExprContext) LambdaConnect() antlr.TerminalNode {
	return s.GetToken(RippletParserLambdaConnect, 0)
}

func (s *FnExprContext) FnBody() IFnBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFnBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFnBodyContext)
}

func (s *FnExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterFnExpr(s)
	}
}

func (s *FnExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitFnExpr(s)
	}
}

func (p *RippletParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *RippletParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, RippletParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeofExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(120)
			p.Match(RippletParserTypeof)
		}
		{
			p.SetState(121)
			p.expression(20)
		}

	case 2:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(122)
			p.Match(RippletParserNot)
		}
		{
			p.SetState(123)
			p.expression(17)
		}

	case 3:
		localctx = NewCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(124)
			p.Identifer()
		}
		{
			p.SetState(125)
			p.Arguments()
		}

	case 4:
		localctx = NewFnExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(127)
			p.FormalParams()
		}
		{
			p.SetState(128)
			p.Match(RippletParserLambdaConnect)
		}
		{
			p.SetState(129)
			p.FnBody()
		}

	case 5:
		localctx = NewOkExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(131)
			p.Match(RippletParserOk)
		}
		{
			p.SetState(132)
			p.expression(9)
		}

	case 6:
		localctx = NewErrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(133)
			p.Match(RippletParserErr)
		}
		{
			p.SetState(134)
			p.expression(8)
		}

	case 7:
		localctx = NewThisExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(135)
			p.Match(RippletParserThis)
		}

	case 8:
		localctx = NewIdentifierExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(136)
			p.Identifer()
		}

	case 9:
		localctx = NewArrayExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(137)
			p.ArrayLiteral()
		}

	case 10:
		localctx = NewObjectExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(138)
			p.ObjectLiteral()
		}

	case 11:
		localctx = NewLiteralExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(139)
			p.Literal()
		}

	case 12:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(140)
			p.Match(RippletParserParenOpen)
		}
		{
			p.SetState(141)
			p.expression(0)
		}
		{
			p.SetState(142)
			p.Match(RippletParserParenClose)
		}

	case 13:
		localctx = NewVoidExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(144)
			p.Match(RippletParserParenOpen)
		}
		{
			p.SetState(145)
			p.Match(RippletParserParenClose)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(177)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEqualityIsExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RippletParserRULE_expression)
				p.SetState(148)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(149)
					p.Match(RippletParserIs)
				}
				{
					p.SetState(150)
					p.expression(20)
				}

			case 2:
				localctx = NewEqualityIsNotExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RippletParserRULE_expression)
				p.SetState(151)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(152)
					p.Match(RippletParserIsNot)
				}
				{
					p.SetState(153)
					p.expression(19)
				}

			case 3:
				localctx = NewPowerExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RippletParserRULE_expression)
				p.SetState(154)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(155)
					p.Match(RippletParserPower)
				}
				{
					p.SetState(156)
					p.expression(14)
				}

			case 4:
				localctx = NewMulExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RippletParserRULE_expression)
				p.SetState(157)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(158)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(RippletParserMultiply-37))|(1<<(RippletParserDivide-37))|(1<<(RippletParserModulus-37)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(159)
					p.expression(14)
				}

			case 5:
				localctx = NewAddExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RippletParserRULE_expression)
				p.SetState(160)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(161)
					_la = p.GetTokenStream().LA(1)

					if !(_la == RippletParserPlus || _la == RippletParserMinus) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(162)
					p.expression(13)
				}

			case 6:
				localctx = NewLogicAndExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RippletParserRULE_expression)
				p.SetState(163)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(164)
					p.Match(RippletParserAnd)
				}
				{
					p.SetState(165)
					p.expression(12)
				}

			case 7:
				localctx = NewLogicOrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RippletParserRULE_expression)
				p.SetState(166)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(167)
					p.Match(RippletParserOr)
				}
				{
					p.SetState(168)
					p.expression(11)
				}

			case 8:
				localctx = NewMemterIdxExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RippletParserRULE_expression)
				p.SetState(169)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(170)
					p.Match(RippletParserBracketOpen)
				}
				{
					p.SetState(171)
					p.expression(0)
				}
				{
					p.SetState(172)
					p.Match(RippletParserBracketClose)
				}

			case 9:
				localctx = NewMemberDotExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RippletParserRULE_expression)
				p.SetState(174)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				{
					p.SetState(175)
					p.Match(RippletParserDot)
				}
				{
					p.SetState(176)
					p.Match(RippletParserIdentifier)
				}

			}

		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IObjDeclareStmtContext is an interface to support dynamic dispatch.
type IObjDeclareStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjDeclareStmtContext differentiates from other interfaces.
	IsObjDeclareStmtContext()
}

type ObjDeclareStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjDeclareStmtContext() *ObjDeclareStmtContext {
	var p = new(ObjDeclareStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_objDeclareStmt
	return p
}

func (*ObjDeclareStmtContext) IsObjDeclareStmtContext() {}

func NewObjDeclareStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjDeclareStmtContext {
	var p = new(ObjDeclareStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_objDeclareStmt

	return p
}

func (s *ObjDeclareStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjDeclareStmtContext) Object() antlr.TerminalNode {
	return s.GetToken(RippletParserObject, 0)
}

func (s *ObjDeclareStmtContext) BraceOpen() antlr.TerminalNode {
	return s.GetToken(RippletParserBraceOpen, 0)
}

func (s *ObjDeclareStmtContext) BraceClose() antlr.TerminalNode {
	return s.GetToken(RippletParserBraceClose, 0)
}

func (s *ObjDeclareStmtContext) ObjProps() IObjPropsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjPropsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjPropsContext)
}

func (s *ObjDeclareStmtContext) ObjMethods() IObjMethodsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjMethodsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjMethodsContext)
}

func (s *ObjDeclareStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjDeclareStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjDeclareStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterObjDeclareStmt(s)
	}
}

func (s *ObjDeclareStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitObjDeclareStmt(s)
	}
}

func (p *RippletParser) ObjDeclareStmt() (localctx IObjDeclareStmtContext) {
	localctx = NewObjDeclareStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RippletParserRULE_objDeclareStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(RippletParserObject)
	}
	{
		p.SetState(183)
		p.Match(RippletParserBraceOpen)
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(184)
			p.ObjProps()
		}

	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RippletParserIdentifier {
		{
			p.SetState(187)
			p.ObjMethods()
		}

	}
	{
		p.SetState(190)
		p.Match(RippletParserBraceClose)
	}

	return localctx
}

// IObjPropsContext is an interface to support dynamic dispatch.
type IObjPropsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjPropsContext differentiates from other interfaces.
	IsObjPropsContext()
}

type ObjPropsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjPropsContext() *ObjPropsContext {
	var p = new(ObjPropsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_objProps
	return p
}

func (*ObjPropsContext) IsObjPropsContext() {}

func NewObjPropsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjPropsContext {
	var p = new(ObjPropsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_objProps

	return p
}

func (s *ObjPropsContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjPropsContext) AllObjProp() []IObjPropContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjPropContext)(nil)).Elem())
	var tst = make([]IObjPropContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjPropContext)
		}
	}

	return tst
}

func (s *ObjPropsContext) ObjProp(i int) IObjPropContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjPropContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjPropContext)
}

func (s *ObjPropsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(RippletParserComma)
}

func (s *ObjPropsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(RippletParserComma, i)
}

func (s *ObjPropsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjPropsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjPropsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterObjProps(s)
	}
}

func (s *ObjPropsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitObjProps(s)
	}
}

func (p *RippletParser) ObjProps() (localctx IObjPropsContext) {
	localctx = NewObjPropsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RippletParserRULE_objProps)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.ObjProp()
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RippletParserComma {
		{
			p.SetState(193)
			p.Match(RippletParserComma)
		}
		{
			p.SetState(194)
			p.ObjProp()
		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IObjPropContext is an interface to support dynamic dispatch.
type IObjPropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjPropContext differentiates from other interfaces.
	IsObjPropContext()
}

type ObjPropContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjPropContext() *ObjPropContext {
	var p = new(ObjPropContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_objProp
	return p
}

func (*ObjPropContext) IsObjPropContext() {}

func NewObjPropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjPropContext {
	var p = new(ObjPropContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_objProp

	return p
}

func (s *ObjPropContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjPropContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RippletParserIdentifier, 0)
}

func (s *ObjPropContext) Assign() antlr.TerminalNode {
	return s.GetToken(RippletParserAssign, 0)
}

func (s *ObjPropContext) ObjPropInit() IObjPropInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjPropInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjPropInitContext)
}

func (s *ObjPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjPropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterObjProp(s)
	}
}

func (s *ObjPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitObjProp(s)
	}
}

func (p *RippletParser) ObjProp() (localctx IObjPropContext) {
	localctx = NewObjPropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RippletParserRULE_objProp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(RippletParserIdentifier)
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RippletParserAssign {
		{
			p.SetState(201)
			p.Match(RippletParserAssign)
		}
		{
			p.SetState(202)
			p.ObjPropInit()
		}

	}

	return localctx
}

// IObjPropInitContext is an interface to support dynamic dispatch.
type IObjPropInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjPropInitContext differentiates from other interfaces.
	IsObjPropInitContext()
}

type ObjPropInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjPropInitContext() *ObjPropInitContext {
	var p = new(ObjPropInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_objPropInit
	return p
}

func (*ObjPropInitContext) IsObjPropInitContext() {}

func NewObjPropInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjPropInitContext {
	var p = new(ObjPropInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_objPropInit

	return p
}

func (s *ObjPropInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjPropInitContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ObjPropInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjPropInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjPropInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterObjPropInit(s)
	}
}

func (s *ObjPropInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitObjPropInit(s)
	}
}

func (p *RippletParser) ObjPropInit() (localctx IObjPropInitContext) {
	localctx = NewObjPropInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RippletParserRULE_objPropInit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.expression(0)
	}

	return localctx
}

// IObjMethodsContext is an interface to support dynamic dispatch.
type IObjMethodsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjMethodsContext differentiates from other interfaces.
	IsObjMethodsContext()
}

type ObjMethodsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjMethodsContext() *ObjMethodsContext {
	var p = new(ObjMethodsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_objMethods
	return p
}

func (*ObjMethodsContext) IsObjMethodsContext() {}

func NewObjMethodsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjMethodsContext {
	var p = new(ObjMethodsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_objMethods

	return p
}

func (s *ObjMethodsContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjMethodsContext) AllObjMethod() []IObjMethodContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjMethodContext)(nil)).Elem())
	var tst = make([]IObjMethodContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjMethodContext)
		}
	}

	return tst
}

func (s *ObjMethodsContext) ObjMethod(i int) IObjMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjMethodContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjMethodContext)
}

func (s *ObjMethodsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjMethodsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjMethodsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterObjMethods(s)
	}
}

func (s *ObjMethodsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitObjMethods(s)
	}
}

func (p *RippletParser) ObjMethods() (localctx IObjMethodsContext) {
	localctx = NewObjMethodsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RippletParserRULE_objMethods)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RippletParserIdentifier {
		{
			p.SetState(207)
			p.ObjMethod()
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IObjMethodContext is an interface to support dynamic dispatch.
type IObjMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjMethodContext differentiates from other interfaces.
	IsObjMethodContext()
}

type ObjMethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjMethodContext() *ObjMethodContext {
	var p = new(ObjMethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_objMethod
	return p
}

func (*ObjMethodContext) IsObjMethodContext() {}

func NewObjMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjMethodContext {
	var p = new(ObjMethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_objMethod

	return p
}

func (s *ObjMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjMethodContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RippletParserIdentifier, 0)
}

func (s *ObjMethodContext) FormalParams() IFormalParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParamsContext)
}

func (s *ObjMethodContext) BlockStmt() IBlockStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockStmtContext)
}

func (s *ObjMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterObjMethod(s)
	}
}

func (s *ObjMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitObjMethod(s)
	}
}

func (p *RippletParser) ObjMethod() (localctx IObjMethodContext) {
	localctx = NewObjMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RippletParserRULE_objMethod)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(RippletParserIdentifier)
	}
	{
		p.SetState(213)
		p.FormalParams()
	}
	{
		p.SetState(214)
		p.BlockStmt()
	}

	return localctx
}

// IAssignStmtContext is an interface to support dynamic dispatch.
type IAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignStmtContext differentiates from other interfaces.
	IsAssignStmtContext()
}

type AssignStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignStmtContext() *AssignStmtContext {
	var p = new(AssignStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_assignStmt
	return p
}

func (*AssignStmtContext) IsAssignStmtContext() {}

func NewAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStmtContext {
	var p = new(AssignStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_assignStmt

	return p
}

func (s *AssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStmtContext) Identifer() IIdentiferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentiferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentiferContext)
}

func (s *AssignStmtContext) Assign() antlr.TerminalNode {
	return s.GetToken(RippletParserAssign, 0)
}

func (s *AssignStmtContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterAssignStmt(s)
	}
}

func (s *AssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitAssignStmt(s)
	}
}

func (p *RippletParser) AssignStmt() (localctx IAssignStmtContext) {
	localctx = NewAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RippletParserRULE_assignStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Identifer()
	}
	{
		p.SetState(217)
		p.Match(RippletParserAssign)
	}
	{
		p.SetState(218)
		p.Statement()
	}

	return localctx
}

// IVarDeclareStmtContext is an interface to support dynamic dispatch.
type IVarDeclareStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclareStmtContext differentiates from other interfaces.
	IsVarDeclareStmtContext()
}

type VarDeclareStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclareStmtContext() *VarDeclareStmtContext {
	var p = new(VarDeclareStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_varDeclareStmt
	return p
}

func (*VarDeclareStmtContext) IsVarDeclareStmtContext() {}

func NewVarDeclareStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclareStmtContext {
	var p = new(VarDeclareStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_varDeclareStmt

	return p
}

func (s *VarDeclareStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclareStmtContext) VarDeclareLhs() IVarDeclareLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclareLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclareLhsContext)
}

func (s *VarDeclareStmtContext) Declare() antlr.TerminalNode {
	return s.GetToken(RippletParserDeclare, 0)
}

func (s *VarDeclareStmtContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *VarDeclareStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclareStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclareStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterVarDeclareStmt(s)
	}
}

func (s *VarDeclareStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitVarDeclareStmt(s)
	}
}

func (p *RippletParser) VarDeclareStmt() (localctx IVarDeclareStmtContext) {
	localctx = NewVarDeclareStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RippletParserRULE_varDeclareStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.VarDeclareLhs()
	}
	{
		p.SetState(221)
		p.Match(RippletParserDeclare)
	}
	{
		p.SetState(222)
		p.Statement()
	}

	return localctx
}

// IVarDeclareLhsContext is an interface to support dynamic dispatch.
type IVarDeclareLhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclareLhsContext differentiates from other interfaces.
	IsVarDeclareLhsContext()
}

type VarDeclareLhsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclareLhsContext() *VarDeclareLhsContext {
	var p = new(VarDeclareLhsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_varDeclareLhs
	return p
}

func (*VarDeclareLhsContext) IsVarDeclareLhsContext() {}

func NewVarDeclareLhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclareLhsContext {
	var p = new(VarDeclareLhsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_varDeclareLhs

	return p
}

func (s *VarDeclareLhsContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclareLhsContext) Identifer() IIdentiferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentiferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentiferContext)
}

func (s *VarDeclareLhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclareLhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclareLhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterVarDeclareLhs(s)
	}
}

func (s *VarDeclareLhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitVarDeclareLhs(s)
	}
}

func (p *RippletParser) VarDeclareLhs() (localctx IVarDeclareLhsContext) {
	localctx = NewVarDeclareLhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RippletParserRULE_varDeclareLhs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Identifer()
	}

	return localctx
}

// IExprStmtContext is an interface to support dynamic dispatch.
type IExprStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprStmtContext differentiates from other interfaces.
	IsExprStmtContext()
}

type ExprStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprStmtContext() *ExprStmtContext {
	var p = new(ExprStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_exprStmt
	return p
}

func (*ExprStmtContext) IsExprStmtContext() {}

func NewExprStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprStmtContext {
	var p = new(ExprStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_exprStmt

	return p
}

func (s *ExprStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterExprStmt(s)
	}
}

func (s *ExprStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitExprStmt(s)
	}
}

func (p *RippletParser) ExprStmt() (localctx IExprStmtContext) {
	localctx = NewExprStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RippletParserRULE_exprStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.expression(0)
	}

	return localctx
}

// IMatchStmtContext is an interface to support dynamic dispatch.
type IMatchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMatchStmtContext differentiates from other interfaces.
	IsMatchStmtContext()
}

type MatchStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchStmtContext() *MatchStmtContext {
	var p = new(MatchStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_matchStmt
	return p
}

func (*MatchStmtContext) IsMatchStmtContext() {}

func NewMatchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchStmtContext {
	var p = new(MatchStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_matchStmt

	return p
}

func (s *MatchStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchStmtContext) Match() antlr.TerminalNode {
	return s.GetToken(RippletParserMatch, 0)
}

func (s *MatchStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MatchStmtContext) BraceOpen() antlr.TerminalNode {
	return s.GetToken(RippletParserBraceOpen, 0)
}

func (s *MatchStmtContext) MathClauses() IMathClausesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathClausesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathClausesContext)
}

func (s *MatchStmtContext) BraceClose() antlr.TerminalNode {
	return s.GetToken(RippletParserBraceClose, 0)
}

func (s *MatchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterMatchStmt(s)
	}
}

func (s *MatchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitMatchStmt(s)
	}
}

func (p *RippletParser) MatchStmt() (localctx IMatchStmtContext) {
	localctx = NewMatchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RippletParserRULE_matchStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(RippletParserMatch)
	}
	{
		p.SetState(229)
		p.expression(0)
	}
	{
		p.SetState(230)
		p.Match(RippletParserBraceOpen)
	}
	{
		p.SetState(231)
		p.MathClauses()
	}
	{
		p.SetState(232)
		p.Match(RippletParserBraceClose)
	}

	return localctx
}

// IMathClausesContext is an interface to support dynamic dispatch.
type IMathClausesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathClausesContext differentiates from other interfaces.
	IsMathClausesContext()
}

type MathClausesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathClausesContext() *MathClausesContext {
	var p = new(MathClausesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_mathClauses
	return p
}

func (*MathClausesContext) IsMathClausesContext() {}

func NewMathClausesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathClausesContext {
	var p = new(MathClausesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_mathClauses

	return p
}

func (s *MathClausesContext) GetParser() antlr.Parser { return s.parser }

func (s *MathClausesContext) AllMathClause() []IMathClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMathClauseContext)(nil)).Elem())
	var tst = make([]IMathClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMathClauseContext)
		}
	}

	return tst
}

func (s *MathClausesContext) MathClause(i int) IMathClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMathClauseContext)
}

func (s *MathClausesContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(RippletParserComma)
}

func (s *MathClausesContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(RippletParserComma, i)
}

func (s *MathClausesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathClausesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathClausesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterMathClauses(s)
	}
}

func (s *MathClausesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitMathClauses(s)
	}
}

func (p *RippletParser) MathClauses() (localctx IMathClausesContext) {
	localctx = NewMathClausesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RippletParserRULE_mathClauses)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.MathClause()
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RippletParserComma {
		{
			p.SetState(235)
			p.Match(RippletParserComma)
		}
		{
			p.SetState(236)
			p.MathClause()
		}

		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMathClauseContext is an interface to support dynamic dispatch.
type IMathClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathClauseContext differentiates from other interfaces.
	IsMathClauseContext()
}

type MathClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathClauseContext() *MathClauseContext {
	var p = new(MathClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_mathClause
	return p
}

func (*MathClauseContext) IsMathClauseContext() {}

func NewMathClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathClauseContext {
	var p = new(MathClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_mathClause

	return p
}

func (s *MathClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *MathClauseContext) MatchClauseTest() IMatchClauseTestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatchClauseTestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatchClauseTestContext)
}

func (s *MathClauseContext) LambdaConnect() antlr.TerminalNode {
	return s.GetToken(RippletParserLambdaConnect, 0)
}

func (s *MathClauseContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *MathClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterMathClause(s)
	}
}

func (s *MathClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitMathClause(s)
	}
}

func (p *RippletParser) MathClause() (localctx IMathClauseContext) {
	localctx = NewMathClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RippletParserRULE_mathClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.MatchClauseTest()
	}
	{
		p.SetState(243)
		p.Match(RippletParserLambdaConnect)
	}
	{
		p.SetState(244)
		p.Statement()
	}

	return localctx
}

// IMatchClauseTestContext is an interface to support dynamic dispatch.
type IMatchClauseTestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMatchClauseTestContext differentiates from other interfaces.
	IsMatchClauseTestContext()
}

type MatchClauseTestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchClauseTestContext() *MatchClauseTestContext {
	var p = new(MatchClauseTestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_matchClauseTest
	return p
}

func (*MatchClauseTestContext) IsMatchClauseTestContext() {}

func NewMatchClauseTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchClauseTestContext {
	var p = new(MatchClauseTestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_matchClauseTest

	return p
}

func (s *MatchClauseTestContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchClauseTestContext) MatchClauseTestVal() IMatchClauseTestValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatchClauseTestValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatchClauseTestValContext)
}

func (s *MatchClauseTestContext) MatchClauseTestUnwrap() IMatchClauseTestUnwrapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatchClauseTestUnwrapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatchClauseTestUnwrapContext)
}

func (s *MatchClauseTestContext) Discard() antlr.TerminalNode {
	return s.GetToken(RippletParserDiscard, 0)
}

func (s *MatchClauseTestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchClauseTestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchClauseTestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterMatchClauseTest(s)
	}
}

func (s *MatchClauseTestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitMatchClauseTest(s)
	}
}

func (p *RippletParser) MatchClauseTest() (localctx IMatchClauseTestContext) {
	localctx = NewMatchClauseTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RippletParserRULE_matchClauseTest)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(249)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RippletParserIntLiteral, RippletParserHexLiteral, RippletParserRealLiteral, RippletParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(246)
			p.MatchClauseTestVal()
		}

	case RippletParserOk, RippletParserErr:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(247)
			p.MatchClauseTestUnwrap()
		}

	case RippletParserDiscard:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(248)
			p.Match(RippletParserDiscard)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMatchClauseTestValContext is an interface to support dynamic dispatch.
type IMatchClauseTestValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMatchClauseTestValContext differentiates from other interfaces.
	IsMatchClauseTestValContext()
}

type MatchClauseTestValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchClauseTestValContext() *MatchClauseTestValContext {
	var p = new(MatchClauseTestValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_matchClauseTestVal
	return p
}

func (*MatchClauseTestValContext) IsMatchClauseTestValContext() {}

func NewMatchClauseTestValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchClauseTestValContext {
	var p = new(MatchClauseTestValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_matchClauseTestVal

	return p
}

func (s *MatchClauseTestValContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchClauseTestValContext) AllMatchClauseVal() []IMatchClauseValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMatchClauseValContext)(nil)).Elem())
	var tst = make([]IMatchClauseValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMatchClauseValContext)
		}
	}

	return tst
}

func (s *MatchClauseTestValContext) MatchClauseVal(i int) IMatchClauseValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatchClauseValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMatchClauseValContext)
}

func (s *MatchClauseTestValContext) AllOr() []antlr.TerminalNode {
	return s.GetTokens(RippletParserOr)
}

func (s *MatchClauseTestValContext) Or(i int) antlr.TerminalNode {
	return s.GetToken(RippletParserOr, i)
}

func (s *MatchClauseTestValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchClauseTestValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchClauseTestValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterMatchClauseTestVal(s)
	}
}

func (s *MatchClauseTestValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitMatchClauseTestVal(s)
	}
}

func (p *RippletParser) MatchClauseTestVal() (localctx IMatchClauseTestValContext) {
	localctx = NewMatchClauseTestValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RippletParserRULE_matchClauseTestVal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.MatchClauseVal()
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RippletParserOr {
		{
			p.SetState(252)
			p.Match(RippletParserOr)
		}
		{
			p.SetState(253)
			p.MatchClauseVal()
		}

		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMatchClauseTestUnwrapContext is an interface to support dynamic dispatch.
type IMatchClauseTestUnwrapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMatchClauseTestUnwrapContext differentiates from other interfaces.
	IsMatchClauseTestUnwrapContext()
}

type MatchClauseTestUnwrapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchClauseTestUnwrapContext() *MatchClauseTestUnwrapContext {
	var p = new(MatchClauseTestUnwrapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_matchClauseTestUnwrap
	return p
}

func (*MatchClauseTestUnwrapContext) IsMatchClauseTestUnwrapContext() {}

func NewMatchClauseTestUnwrapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchClauseTestUnwrapContext {
	var p = new(MatchClauseTestUnwrapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_matchClauseTestUnwrap

	return p
}

func (s *MatchClauseTestUnwrapContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchClauseTestUnwrapContext) ParenOpen() antlr.TerminalNode {
	return s.GetToken(RippletParserParenOpen, 0)
}

func (s *MatchClauseTestUnwrapContext) Identifer() IIdentiferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentiferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentiferContext)
}

func (s *MatchClauseTestUnwrapContext) ParenClose() antlr.TerminalNode {
	return s.GetToken(RippletParserParenClose, 0)
}

func (s *MatchClauseTestUnwrapContext) Ok() antlr.TerminalNode {
	return s.GetToken(RippletParserOk, 0)
}

func (s *MatchClauseTestUnwrapContext) Err() antlr.TerminalNode {
	return s.GetToken(RippletParserErr, 0)
}

func (s *MatchClauseTestUnwrapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchClauseTestUnwrapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchClauseTestUnwrapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterMatchClauseTestUnwrap(s)
	}
}

func (s *MatchClauseTestUnwrapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitMatchClauseTestUnwrap(s)
	}
}

func (p *RippletParser) MatchClauseTestUnwrap() (localctx IMatchClauseTestUnwrapContext) {
	localctx = NewMatchClauseTestUnwrapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RippletParserRULE_matchClauseTestUnwrap)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RippletParserOk || _la == RippletParserErr) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(260)
		p.Match(RippletParserParenOpen)
	}
	{
		p.SetState(261)
		p.Identifer()
	}
	{
		p.SetState(262)
		p.Match(RippletParserParenClose)
	}

	return localctx
}

// IMatchClauseValContext is an interface to support dynamic dispatch.
type IMatchClauseValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMatchClauseValContext differentiates from other interfaces.
	IsMatchClauseValContext()
}

type MatchClauseValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchClauseValContext() *MatchClauseValContext {
	var p = new(MatchClauseValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_matchClauseVal
	return p
}

func (*MatchClauseValContext) IsMatchClauseValContext() {}

func NewMatchClauseValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchClauseValContext {
	var p = new(MatchClauseValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_matchClauseVal

	return p
}

func (s *MatchClauseValContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchClauseValContext) Identifer() IIdentiferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentiferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentiferContext)
}

func (s *MatchClauseValContext) NumberLiteral() INumberLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberLiteralContext)
}

func (s *MatchClauseValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchClauseValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchClauseValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterMatchClauseVal(s)
	}
}

func (s *MatchClauseValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitMatchClauseVal(s)
	}
}

func (p *RippletParser) MatchClauseVal() (localctx IMatchClauseValContext) {
	localctx = NewMatchClauseValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RippletParserRULE_matchClauseVal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(266)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RippletParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.Identifer()
		}

	case RippletParserIntLiteral, RippletParserHexLiteral, RippletParserRealLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(265)
			p.NumberLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFnDeclareStmtContext is an interface to support dynamic dispatch.
type IFnDeclareStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFnDeclareStmtContext differentiates from other interfaces.
	IsFnDeclareStmtContext()
}

type FnDeclareStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnDeclareStmtContext() *FnDeclareStmtContext {
	var p = new(FnDeclareStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_fnDeclareStmt
	return p
}

func (*FnDeclareStmtContext) IsFnDeclareStmtContext() {}

func NewFnDeclareStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnDeclareStmtContext {
	var p = new(FnDeclareStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_fnDeclareStmt

	return p
}

func (s *FnDeclareStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FnDeclareStmtContext) Fn() antlr.TerminalNode {
	return s.GetToken(RippletParserFn, 0)
}

func (s *FnDeclareStmtContext) FnName() IFnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFnNameContext)
}

func (s *FnDeclareStmtContext) FormalParams() IFormalParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParamsContext)
}

func (s *FnDeclareStmtContext) FnBody() IFnBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFnBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFnBodyContext)
}

func (s *FnDeclareStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnDeclareStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnDeclareStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterFnDeclareStmt(s)
	}
}

func (s *FnDeclareStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitFnDeclareStmt(s)
	}
}

func (p *RippletParser) FnDeclareStmt() (localctx IFnDeclareStmtContext) {
	localctx = NewFnDeclareStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RippletParserRULE_fnDeclareStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(RippletParserFn)
	}
	{
		p.SetState(269)
		p.FnName()
	}
	{
		p.SetState(270)
		p.FormalParams()
	}
	{
		p.SetState(271)
		p.FnBody()
	}

	return localctx
}

// IFnNameContext is an interface to support dynamic dispatch.
type IFnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFnNameContext differentiates from other interfaces.
	IsFnNameContext()
}

type FnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnNameContext() *FnNameContext {
	var p = new(FnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_fnName
	return p
}

func (*FnNameContext) IsFnNameContext() {}

func NewFnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnNameContext {
	var p = new(FnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_fnName

	return p
}

func (s *FnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FnNameContext) Identifer() IIdentiferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentiferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentiferContext)
}

func (s *FnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterFnName(s)
	}
}

func (s *FnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitFnName(s)
	}
}

func (p *RippletParser) FnName() (localctx IFnNameContext) {
	localctx = NewFnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RippletParserRULE_fnName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Identifer()
	}

	return localctx
}

// IIdentiferContext is an interface to support dynamic dispatch.
type IIdentiferContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentiferContext differentiates from other interfaces.
	IsIdentiferContext()
}

type IdentiferContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentiferContext() *IdentiferContext {
	var p = new(IdentiferContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_identifer
	return p
}

func (*IdentiferContext) IsIdentiferContext() {}

func NewIdentiferContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentiferContext {
	var p = new(IdentiferContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_identifer

	return p
}

func (s *IdentiferContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentiferContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RippletParserIdentifier, 0)
}

func (s *IdentiferContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentiferContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentiferContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterIdentifer(s)
	}
}

func (s *IdentiferContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitIdentifer(s)
	}
}

func (p *RippletParser) Identifer() (localctx IIdentiferContext) {
	localctx = NewIdentiferContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, RippletParserRULE_identifer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.Match(RippletParserIdentifier)
	}

	return localctx
}

// IFormalParamsContext is an interface to support dynamic dispatch.
type IFormalParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParamsContext differentiates from other interfaces.
	IsFormalParamsContext()
}

type FormalParamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParamsContext() *FormalParamsContext {
	var p = new(FormalParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_formalParams
	return p
}

func (*FormalParamsContext) IsFormalParamsContext() {}

func NewFormalParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParamsContext {
	var p = new(FormalParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_formalParams

	return p
}

func (s *FormalParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParamsContext) ParenOpen() antlr.TerminalNode {
	return s.GetToken(RippletParserParenOpen, 0)
}

func (s *FormalParamsContext) ParenClose() antlr.TerminalNode {
	return s.GetToken(RippletParserParenClose, 0)
}

func (s *FormalParamsContext) FormalParamList() IFormalParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParamListContext)
}

func (s *FormalParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterFormalParams(s)
	}
}

func (s *FormalParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitFormalParams(s)
	}
}

func (p *RippletParser) FormalParams() (localctx IFormalParamsContext) {
	localctx = NewFormalParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, RippletParserRULE_formalParams)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Match(RippletParserParenOpen)
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RippletParserEllipsis || _la == RippletParserIdentifier {
		{
			p.SetState(278)
			p.FormalParamList()
		}

	}
	{
		p.SetState(281)
		p.Match(RippletParserParenClose)
	}

	return localctx
}

// IFormalParamListContext is an interface to support dynamic dispatch.
type IFormalParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParamListContext differentiates from other interfaces.
	IsFormalParamListContext()
}

type FormalParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParamListContext() *FormalParamListContext {
	var p = new(FormalParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_formalParamList
	return p
}

func (*FormalParamListContext) IsFormalParamListContext() {}

func NewFormalParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParamListContext {
	var p = new(FormalParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_formalParamList

	return p
}

func (s *FormalParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParamListContext) AllIdentifer() []IIdentiferContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentiferContext)(nil)).Elem())
	var tst = make([]IIdentiferContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentiferContext)
		}
	}

	return tst
}

func (s *FormalParamListContext) Identifer(i int) IIdentiferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentiferContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentiferContext)
}

func (s *FormalParamListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(RippletParserComma)
}

func (s *FormalParamListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(RippletParserComma, i)
}

func (s *FormalParamListContext) RestParamArg() IRestParamArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRestParamArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRestParamArgContext)
}

func (s *FormalParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterFormalParamList(s)
	}
}

func (s *FormalParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitFormalParamList(s)
	}
}

func (p *RippletParser) FormalParamList() (localctx IFormalParamListContext) {
	localctx = NewFormalParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, RippletParserRULE_formalParamList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(296)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RippletParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(283)
			p.Identifer()
		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(284)
					p.Match(RippletParserComma)
				}
				{
					p.SetState(285)
					p.Identifer()
				}

			}
			p.SetState(290)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
		}
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RippletParserComma {
			{
				p.SetState(291)
				p.Match(RippletParserComma)
			}
			{
				p.SetState(292)
				p.RestParamArg()
			}

		}

	case RippletParserEllipsis:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.RestParamArg()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFnBodyContext is an interface to support dynamic dispatch.
type IFnBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFnBodyContext differentiates from other interfaces.
	IsFnBodyContext()
}

type FnBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnBodyContext() *FnBodyContext {
	var p = new(FnBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_fnBody
	return p
}

func (*FnBodyContext) IsFnBodyContext() {}

func NewFnBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnBodyContext {
	var p = new(FnBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_fnBody

	return p
}

func (s *FnBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FnBodyContext) BraceOpen() antlr.TerminalNode {
	return s.GetToken(RippletParserBraceOpen, 0)
}

func (s *FnBodyContext) BraceClose() antlr.TerminalNode {
	return s.GetToken(RippletParserBraceClose, 0)
}

func (s *FnBodyContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *FnBodyContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *FnBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterFnBody(s)
	}
}

func (s *FnBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitFnBody(s)
	}
}

func (p *RippletParser) FnBody() (localctx IFnBodyContext) {
	localctx = NewFnBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, RippletParserRULE_fnBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Match(RippletParserBraceOpen)
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RippletParserIf)|(1<<RippletParserMatch)|(1<<RippletParserFn)|(1<<RippletParserRepeat)|(1<<RippletParserTypeof)|(1<<RippletParserNot)|(1<<RippletParserOk)|(1<<RippletParserErr)|(1<<RippletParserObject)|(1<<RippletParserThis)|(1<<RippletParserRegexpStart)|(1<<RippletParserBraceOpen)|(1<<RippletParserBracketOpen)|(1<<RippletParserParenOpen))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(RippletParserNilLiteral-51))|(1<<(RippletParserBooLiteral-51))|(1<<(RippletParserIntLiteral-51))|(1<<(RippletParserHexLiteral-51))|(1<<(RippletParserRealLiteral-51))|(1<<(RippletParserIdentifier-51))|(1<<(RippletParserStringOpen-51)))) != 0) {
		{
			p.SetState(299)
			p.Statement()
		}

		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(305)
		p.Match(RippletParserBraceClose)
	}

	return localctx
}

// IRestParamArgContext is an interface to support dynamic dispatch.
type IRestParamArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRestParamArgContext differentiates from other interfaces.
	IsRestParamArgContext()
}

type RestParamArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRestParamArgContext() *RestParamArgContext {
	var p = new(RestParamArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_restParamArg
	return p
}

func (*RestParamArgContext) IsRestParamArgContext() {}

func NewRestParamArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RestParamArgContext {
	var p = new(RestParamArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_restParamArg

	return p
}

func (s *RestParamArgContext) GetParser() antlr.Parser { return s.parser }

func (s *RestParamArgContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(RippletParserEllipsis, 0)
}

func (s *RestParamArgContext) Identifer() IIdentiferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentiferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentiferContext)
}

func (s *RestParamArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RestParamArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RestParamArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterRestParamArg(s)
	}
}

func (s *RestParamArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitRestParamArg(s)
	}
}

func (p *RippletParser) RestParamArg() (localctx IRestParamArgContext) {
	localctx = NewRestParamArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, RippletParserRULE_restParamArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(RippletParserEllipsis)
	}
	{
		p.SetState(308)
		p.Identifer()
	}

	return localctx
}

// IBlockStmtContext is an interface to support dynamic dispatch.
type IBlockStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStmtContext differentiates from other interfaces.
	IsBlockStmtContext()
}

type BlockStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStmtContext() *BlockStmtContext {
	var p = new(BlockStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_blockStmt
	return p
}

func (*BlockStmtContext) IsBlockStmtContext() {}

func NewBlockStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStmtContext {
	var p = new(BlockStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_blockStmt

	return p
}

func (s *BlockStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStmtContext) BraceOpen() antlr.TerminalNode {
	return s.GetToken(RippletParserBraceOpen, 0)
}

func (s *BlockStmtContext) BraceClose() antlr.TerminalNode {
	return s.GetToken(RippletParserBraceClose, 0)
}

func (s *BlockStmtContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockStmtContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterBlockStmt(s)
	}
}

func (s *BlockStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitBlockStmt(s)
	}
}

func (p *RippletParser) BlockStmt() (localctx IBlockStmtContext) {
	localctx = NewBlockStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, RippletParserRULE_blockStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(RippletParserBraceOpen)
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RippletParserIf)|(1<<RippletParserMatch)|(1<<RippletParserFn)|(1<<RippletParserRepeat)|(1<<RippletParserTypeof)|(1<<RippletParserNot)|(1<<RippletParserOk)|(1<<RippletParserErr)|(1<<RippletParserObject)|(1<<RippletParserThis)|(1<<RippletParserRegexpStart)|(1<<RippletParserBraceOpen)|(1<<RippletParserBracketOpen)|(1<<RippletParserParenOpen))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(RippletParserNilLiteral-51))|(1<<(RippletParserBooLiteral-51))|(1<<(RippletParserIntLiteral-51))|(1<<(RippletParserHexLiteral-51))|(1<<(RippletParserRealLiteral-51))|(1<<(RippletParserIdentifier-51))|(1<<(RippletParserStringOpen-51)))) != 0) {
		{
			p.SetState(311)
			p.Statement()
		}

		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(317)
		p.Match(RippletParserBraceClose)
	}

	return localctx
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_ifStmt
	return p
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) If() antlr.TerminalNode {
	return s.GetToken(RippletParserIf, 0)
}

func (s *IfStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfStmtContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStmtContext) Then() antlr.TerminalNode {
	return s.GetToken(RippletParserThen, 0)
}

func (s *IfStmtContext) Else() antlr.TerminalNode {
	return s.GetToken(RippletParserElse, 0)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (p *RippletParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, RippletParserRULE_ifStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(319)
		p.Match(RippletParserIf)
	}
	{
		p.SetState(320)
		p.expression(0)
	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RippletParserThen {
		{
			p.SetState(321)
			p.Match(RippletParserThen)
		}

	}
	{
		p.SetState(324)
		p.Statement()
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(325)
			p.Match(RippletParserElse)
		}
		{
			p.SetState(326)
			p.Statement()
		}

	}

	return localctx
}

// IRepeatStmtContext is an interface to support dynamic dispatch.
type IRepeatStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRepeatStmtContext differentiates from other interfaces.
	IsRepeatStmtContext()
}

type RepeatStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeatStmtContext() *RepeatStmtContext {
	var p = new(RepeatStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_repeatStmt
	return p
}

func (*RepeatStmtContext) IsRepeatStmtContext() {}

func NewRepeatStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatStmtContext {
	var p = new(RepeatStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_repeatStmt

	return p
}

func (s *RepeatStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatStmtContext) Repeat() antlr.TerminalNode {
	return s.GetToken(RippletParserRepeat, 0)
}

func (s *RepeatStmtContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *RepeatStmtContext) Until() antlr.TerminalNode {
	return s.GetToken(RippletParserUntil, 0)
}

func (s *RepeatStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RepeatStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepeatStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterRepeatStmt(s)
	}
}

func (s *RepeatStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitRepeatStmt(s)
	}
}

func (p *RippletParser) RepeatStmt() (localctx IRepeatStmtContext) {
	localctx = NewRepeatStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, RippletParserRULE_repeatStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.Match(RippletParserRepeat)
	}
	{
		p.SetState(330)
		p.Statement()
	}
	{
		p.SetState(331)
		p.Match(RippletParserUntil)
	}
	{
		p.SetState(332)
		p.expression(0)
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) ParenOpen() antlr.TerminalNode {
	return s.GetToken(RippletParserParenOpen, 0)
}

func (s *ArgumentsContext) ParenClose() antlr.TerminalNode {
	return s.GetToken(RippletParserParenClose, 0)
}

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(RippletParserComma)
}

func (s *ArgumentsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(RippletParserComma, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *RippletParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, RippletParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.Match(RippletParserParenOpen)
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RippletParserTypeof)|(1<<RippletParserNot)|(1<<RippletParserOk)|(1<<RippletParserErr)|(1<<RippletParserThis)|(1<<RippletParserRegexpStart)|(1<<RippletParserBraceOpen)|(1<<RippletParserBracketOpen)|(1<<RippletParserParenOpen))) != 0) || (((_la-49)&-(0x1f+1)) == 0 && ((1<<uint((_la-49)))&((1<<(RippletParserEllipsis-49))|(1<<(RippletParserNilLiteral-49))|(1<<(RippletParserBooLiteral-49))|(1<<(RippletParserIntLiteral-49))|(1<<(RippletParserHexLiteral-49))|(1<<(RippletParserRealLiteral-49))|(1<<(RippletParserIdentifier-49))|(1<<(RippletParserStringOpen-49)))) != 0) {
		{
			p.SetState(335)
			p.Argument()
		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(336)
					p.Match(RippletParserComma)
				}
				{
					p.SetState(337)
					p.Argument()
				}

			}
			p.SetState(342)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RippletParserComma {
			{
				p.SetState(343)
				p.Match(RippletParserComma)
			}

		}

	}
	{
		p.SetState(348)
		p.Match(RippletParserParenClose)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RippletParserIdentifier, 0)
}

func (s *ArgumentContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(RippletParserEllipsis, 0)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *RippletParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, RippletParserRULE_argument)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RippletParserEllipsis {
		{
			p.SetState(350)
			p.Match(RippletParserEllipsis)
		}

	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(353)
			p.expression(0)
		}

	case 2:
		{
			p.SetState(354)
			p.Match(RippletParserIdentifier)
		}

	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NilLiteral() INilLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INilLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INilLiteralContext)
}

func (s *LiteralContext) BoolLiteral() IBoolLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolLiteralContext)
}

func (s *LiteralContext) NumberLiteral() INumberLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberLiteralContext)
}

func (s *LiteralContext) RegexpLiteral() IRegexpLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegexpLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegexpLiteralContext)
}

func (s *LiteralContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *RippletParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, RippletParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(362)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RippletParserNilLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(357)
			p.NilLiteral()
		}

	case RippletParserBooLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(358)
			p.BoolLiteral()
		}

	case RippletParserIntLiteral, RippletParserHexLiteral, RippletParserRealLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(359)
			p.NumberLiteral()
		}

	case RippletParserRegexpStart:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(360)
			p.RegexpLiteral()
		}

	case RippletParserStringOpen:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(361)
			p.StringLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRegexpLiteralContext is an interface to support dynamic dispatch.
type IRegexpLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegexpLiteralContext differentiates from other interfaces.
	IsRegexpLiteralContext()
}

type RegexpLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexpLiteralContext() *RegexpLiteralContext {
	var p = new(RegexpLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_regexpLiteral
	return p
}

func (*RegexpLiteralContext) IsRegexpLiteralContext() {}

func NewRegexpLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexpLiteralContext {
	var p = new(RegexpLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_regexpLiteral

	return p
}

func (s *RegexpLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexpLiteralContext) RegexpStart() antlr.TerminalNode {
	return s.GetToken(RippletParserRegexpStart, 0)
}

func (s *RegexpLiteralContext) RexexpClose() antlr.TerminalNode {
	return s.GetToken(RippletParserRexexpClose, 0)
}

func (s *RegexpLiteralContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(RippletParserIdentifier)
}

func (s *RegexpLiteralContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(RippletParserIdentifier, i)
}

func (s *RegexpLiteralContext) AllRegexpContent() []antlr.TerminalNode {
	return s.GetTokens(RippletParserRegexpContent)
}

func (s *RegexpLiteralContext) RegexpContent(i int) antlr.TerminalNode {
	return s.GetToken(RippletParserRegexpContent, i)
}

func (s *RegexpLiteralContext) AllRegexpComment() []antlr.TerminalNode {
	return s.GetTokens(RippletParserRegexpComment)
}

func (s *RegexpLiteralContext) RegexpComment(i int) antlr.TerminalNode {
	return s.GetToken(RippletParserRegexpComment, i)
}

func (s *RegexpLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexpLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexpLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterRegexpLiteral(s)
	}
}

func (s *RegexpLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitRegexpLiteral(s)
	}
}

func (p *RippletParser) RegexpLiteral() (localctx IRegexpLiteralContext) {
	localctx = NewRegexpLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, RippletParserRULE_regexpLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.Match(RippletParserRegexpStart)
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RippletParserRegexpComment || _la == RippletParserRegexpContent {
		{
			p.SetState(365)
			_la = p.GetTokenStream().LA(1)

			if !(_la == RippletParserRegexpComment || _la == RippletParserRegexpContent) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(370)
		p.Match(RippletParserRexexpClose)
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(371)
				p.Match(RippletParserIdentifier)
			}

		}
		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
	}

	return localctx
}

// IBoolLiteralContext is an interface to support dynamic dispatch.
type IBoolLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolLiteralContext differentiates from other interfaces.
	IsBoolLiteralContext()
}

type BoolLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolLiteralContext() *BoolLiteralContext {
	var p = new(BoolLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_boolLiteral
	return p
}

func (*BoolLiteralContext) IsBoolLiteralContext() {}

func NewBoolLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_boolLiteral

	return p
}

func (s *BoolLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolLiteralContext) BooLiteral() antlr.TerminalNode {
	return s.GetToken(RippletParserBooLiteral, 0)
}

func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitBoolLiteral(s)
	}
}

func (p *RippletParser) BoolLiteral() (localctx IBoolLiteralContext) {
	localctx = NewBoolLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, RippletParserRULE_boolLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.Match(RippletParserBooLiteral)
	}

	return localctx
}

// INilLiteralContext is an interface to support dynamic dispatch.
type INilLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNilLiteralContext differentiates from other interfaces.
	IsNilLiteralContext()
}

type NilLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNilLiteralContext() *NilLiteralContext {
	var p = new(NilLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_nilLiteral
	return p
}

func (*NilLiteralContext) IsNilLiteralContext() {}

func NewNilLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NilLiteralContext {
	var p = new(NilLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_nilLiteral

	return p
}

func (s *NilLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NilLiteralContext) NilLiteral() antlr.TerminalNode {
	return s.GetToken(RippletParserNilLiteral, 0)
}

func (s *NilLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NilLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterNilLiteral(s)
	}
}

func (s *NilLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitNilLiteral(s)
	}
}

func (p *RippletParser) NilLiteral() (localctx INilLiteralContext) {
	localctx = NewNilLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, RippletParserRULE_nilLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)
		p.Match(RippletParserNilLiteral)
	}

	return localctx
}

// INumberLiteralContext is an interface to support dynamic dispatch.
type INumberLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberLiteralContext differentiates from other interfaces.
	IsNumberLiteralContext()
}

type NumberLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberLiteralContext() *NumberLiteralContext {
	var p = new(NumberLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_numberLiteral
	return p
}

func (*NumberLiteralContext) IsNumberLiteralContext() {}

func NewNumberLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_numberLiteral

	return p
}

func (s *NumberLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberLiteralContext) IntLiteral() IIntLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntLiteralContext)
}

func (s *NumberLiteralContext) HexLiteral() IHexLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHexLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHexLiteralContext)
}

func (s *NumberLiteralContext) RealLiteral() IRealLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitNumberLiteral(s)
	}
}

func (p *RippletParser) NumberLiteral() (localctx INumberLiteralContext) {
	localctx = NewNumberLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, RippletParserRULE_numberLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(384)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RippletParserIntLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(381)
			p.IntLiteral()
		}

	case RippletParserHexLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(382)
			p.HexLiteral()
		}

	case RippletParserRealLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(383)
			p.RealLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntLiteralContext is an interface to support dynamic dispatch.
type IIntLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntLiteralContext differentiates from other interfaces.
	IsIntLiteralContext()
}

type IntLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntLiteralContext() *IntLiteralContext {
	var p = new(IntLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_intLiteral
	return p
}

func (*IntLiteralContext) IsIntLiteralContext() {}

func NewIntLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntLiteralContext {
	var p = new(IntLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_intLiteral

	return p
}

func (s *IntLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntLiteralContext) IntLiteral() antlr.TerminalNode {
	return s.GetToken(RippletParserIntLiteral, 0)
}

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterIntLiteral(s)
	}
}

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitIntLiteral(s)
	}
}

func (p *RippletParser) IntLiteral() (localctx IIntLiteralContext) {
	localctx = NewIntLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, RippletParserRULE_intLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		p.Match(RippletParserIntLiteral)
	}

	return localctx
}

// IHexLiteralContext is an interface to support dynamic dispatch.
type IHexLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHexLiteralContext differentiates from other interfaces.
	IsHexLiteralContext()
}

type HexLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHexLiteralContext() *HexLiteralContext {
	var p = new(HexLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_hexLiteral
	return p
}

func (*HexLiteralContext) IsHexLiteralContext() {}

func NewHexLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HexLiteralContext {
	var p = new(HexLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_hexLiteral

	return p
}

func (s *HexLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *HexLiteralContext) HexLiteral() antlr.TerminalNode {
	return s.GetToken(RippletParserHexLiteral, 0)
}

func (s *HexLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HexLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterHexLiteral(s)
	}
}

func (s *HexLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitHexLiteral(s)
	}
}

func (p *RippletParser) HexLiteral() (localctx IHexLiteralContext) {
	localctx = NewHexLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, RippletParserRULE_hexLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(388)
		p.Match(RippletParserHexLiteral)
	}

	return localctx
}

// IRealLiteralContext is an interface to support dynamic dispatch.
type IRealLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRealLiteralContext differentiates from other interfaces.
	IsRealLiteralContext()
}

type RealLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealLiteralContext() *RealLiteralContext {
	var p = new(RealLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_realLiteral
	return p
}

func (*RealLiteralContext) IsRealLiteralContext() {}

func NewRealLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealLiteralContext {
	var p = new(RealLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_realLiteral

	return p
}

func (s *RealLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *RealLiteralContext) RealLiteral() antlr.TerminalNode {
	return s.GetToken(RippletParserRealLiteral, 0)
}

func (s *RealLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterRealLiteral(s)
	}
}

func (s *RealLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitRealLiteral(s)
	}
}

func (p *RippletParser) RealLiteral() (localctx IRealLiteralContext) {
	localctx = NewRealLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, RippletParserRULE_realLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(RippletParserRealLiteral)
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) StringOpen() antlr.TerminalNode {
	return s.GetToken(RippletParserStringOpen, 0)
}

func (s *StringLiteralContext) StringClose() antlr.TerminalNode {
	return s.GetToken(RippletParserStringClose, 0)
}

func (s *StringLiteralContext) AllStringQuoted() []antlr.TerminalNode {
	return s.GetTokens(RippletParserStringQuoted)
}

func (s *StringLiteralContext) StringQuoted(i int) antlr.TerminalNode {
	return s.GetToken(RippletParserStringQuoted, i)
}

func (s *StringLiteralContext) AllStringInterp() []IStringInterpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringInterpContext)(nil)).Elem())
	var tst = make([]IStringInterpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringInterpContext)
		}
	}

	return tst
}

func (s *StringLiteralContext) StringInterp(i int) IStringInterpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringInterpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringInterpContext)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (p *RippletParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, RippletParserRULE_stringLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		p.Match(RippletParserStringOpen)
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RippletParserStringInterpStart || _la == RippletParserStringQuoted {
		p.SetState(395)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case RippletParserStringQuoted:
			{
				p.SetState(393)
				p.Match(RippletParserStringQuoted)
			}

		case RippletParserStringInterpStart:
			{
				p.SetState(394)
				p.StringInterp()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(400)
		p.Match(RippletParserStringClose)
	}

	return localctx
}

// IStringQuotedContext is an interface to support dynamic dispatch.
type IStringQuotedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringQuotedContext differentiates from other interfaces.
	IsStringQuotedContext()
}

type StringQuotedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringQuotedContext() *StringQuotedContext {
	var p = new(StringQuotedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_stringQuoted
	return p
}

func (*StringQuotedContext) IsStringQuotedContext() {}

func NewStringQuotedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringQuotedContext {
	var p = new(StringQuotedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_stringQuoted

	return p
}

func (s *StringQuotedContext) GetParser() antlr.Parser { return s.parser }

func (s *StringQuotedContext) StringQuoted() antlr.TerminalNode {
	return s.GetToken(RippletParserStringQuoted, 0)
}

func (s *StringQuotedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringQuotedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringQuotedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterStringQuoted(s)
	}
}

func (s *StringQuotedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitStringQuoted(s)
	}
}

func (p *RippletParser) StringQuoted() (localctx IStringQuotedContext) {
	localctx = NewStringQuotedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, RippletParserRULE_stringQuoted)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		p.Match(RippletParserStringQuoted)
	}

	return localctx
}

// IStringInterpContext is an interface to support dynamic dispatch.
type IStringInterpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringInterpContext differentiates from other interfaces.
	IsStringInterpContext()
}

type StringInterpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringInterpContext() *StringInterpContext {
	var p = new(StringInterpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_stringInterp
	return p
}

func (*StringInterpContext) IsStringInterpContext() {}

func NewStringInterpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringInterpContext {
	var p = new(StringInterpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_stringInterp

	return p
}

func (s *StringInterpContext) GetParser() antlr.Parser { return s.parser }

func (s *StringInterpContext) StringInterpStart() antlr.TerminalNode {
	return s.GetToken(RippletParserStringInterpStart, 0)
}

func (s *StringInterpContext) StringInterpExpr() IStringInterpExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringInterpExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringInterpExprContext)
}

func (s *StringInterpContext) BraceClose() antlr.TerminalNode {
	return s.GetToken(RippletParserBraceClose, 0)
}

func (s *StringInterpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringInterpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringInterpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterStringInterp(s)
	}
}

func (s *StringInterpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitStringInterp(s)
	}
}

func (p *RippletParser) StringInterp() (localctx IStringInterpContext) {
	localctx = NewStringInterpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, RippletParserRULE_stringInterp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(404)
		p.Match(RippletParserStringInterpStart)
	}
	{
		p.SetState(405)
		p.StringInterpExpr()
	}
	{
		p.SetState(406)
		p.Match(RippletParserBraceClose)
	}

	return localctx
}

// IStringInterpExprContext is an interface to support dynamic dispatch.
type IStringInterpExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringInterpExprContext differentiates from other interfaces.
	IsStringInterpExprContext()
}

type StringInterpExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringInterpExprContext() *StringInterpExprContext {
	var p = new(StringInterpExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_stringInterpExpr
	return p
}

func (*StringInterpExprContext) IsStringInterpExprContext() {}

func NewStringInterpExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringInterpExprContext {
	var p = new(StringInterpExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_stringInterpExpr

	return p
}

func (s *StringInterpExprContext) GetParser() antlr.Parser { return s.parser }

func (s *StringInterpExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StringInterpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringInterpExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringInterpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterStringInterpExpr(s)
	}
}

func (s *StringInterpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitStringInterpExpr(s)
	}
}

func (p *RippletParser) StringInterpExpr() (localctx IStringInterpExprContext) {
	localctx = NewStringInterpExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, RippletParserRULE_stringInterpExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(408)
		p.expression(0)
	}

	return localctx
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_arrayLiteral
	return p
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) BracketOpen() antlr.TerminalNode {
	return s.GetToken(RippletParserBracketOpen, 0)
}

func (s *ArrayLiteralContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(RippletParserBracketClose, 0)
}

func (s *ArrayLiteralContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArrayLiteralContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayLiteralContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(RippletParserComma)
}

func (s *ArrayLiteralContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(RippletParserComma, i)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (p *RippletParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, RippletParserRULE_arrayLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(410)
		p.Match(RippletParserBracketOpen)
	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RippletParserTypeof)|(1<<RippletParserNot)|(1<<RippletParserOk)|(1<<RippletParserErr)|(1<<RippletParserThis)|(1<<RippletParserRegexpStart)|(1<<RippletParserBraceOpen)|(1<<RippletParserBracketOpen)|(1<<RippletParserParenOpen))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(RippletParserNilLiteral-51))|(1<<(RippletParserBooLiteral-51))|(1<<(RippletParserIntLiteral-51))|(1<<(RippletParserHexLiteral-51))|(1<<(RippletParserRealLiteral-51))|(1<<(RippletParserIdentifier-51))|(1<<(RippletParserStringOpen-51)))) != 0) {
		{
			p.SetState(411)
			p.expression(0)
		}
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RippletParserComma {
			{
				p.SetState(412)
				p.Match(RippletParserComma)
			}
			{
				p.SetState(413)
				p.expression(0)
			}

			p.SetState(418)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(423)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(424)
		p.Match(RippletParserBracketClose)
	}

	return localctx
}

// IObjectLiteralContext is an interface to support dynamic dispatch.
type IObjectLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectLiteralContext differentiates from other interfaces.
	IsObjectLiteralContext()
}

type ObjectLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectLiteralContext() *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_objectLiteral
	return p
}

func (*ObjectLiteralContext) IsObjectLiteralContext() {}

func NewObjectLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_objectLiteral

	return p
}

func (s *ObjectLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectLiteralContext) BraceOpen() antlr.TerminalNode {
	return s.GetToken(RippletParserBraceOpen, 0)
}

func (s *ObjectLiteralContext) BraceClose() antlr.TerminalNode {
	return s.GetToken(RippletParserBraceClose, 0)
}

func (s *ObjectLiteralContext) AllProp() []IPropContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropContext)(nil)).Elem())
	var tst = make([]IPropContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropContext)
		}
	}

	return tst
}

func (s *ObjectLiteralContext) Prop(i int) IPropContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropContext)
}

func (s *ObjectLiteralContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(RippletParserComma)
}

func (s *ObjectLiteralContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(RippletParserComma, i)
}

func (s *ObjectLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitObjectLiteral(s)
	}
}

func (p *RippletParser) ObjectLiteral() (localctx IObjectLiteralContext) {
	localctx = NewObjectLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, RippletParserRULE_objectLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)
		p.Match(RippletParserBraceOpen)
	}
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(RippletParserBracketOpen-26))|(1<<(RippletParserIntLiteral-26))|(1<<(RippletParserHexLiteral-26))|(1<<(RippletParserRealLiteral-26))|(1<<(RippletParserIdentifier-26))|(1<<(RippletParserStringOpen-26)))) != 0 {
		{
			p.SetState(427)
			p.Prop()
		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(428)
					p.Match(RippletParserComma)
				}
				{
					p.SetState(429)
					p.Prop()
				}

			}
			p.SetState(434)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
		}

	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RippletParserComma {
		{
			p.SetState(437)
			p.Match(RippletParserComma)
		}

	}
	{
		p.SetState(440)
		p.Match(RippletParserBraceClose)
	}

	return localctx
}

// IPropContext is an interface to support dynamic dispatch.
type IPropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropContext differentiates from other interfaces.
	IsPropContext()
}

type PropContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropContext() *PropContext {
	var p = new(PropContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_prop
	return p
}

func (*PropContext) IsPropContext() {}

func NewPropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropContext {
	var p = new(PropContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_prop

	return p
}

func (s *PropContext) GetParser() antlr.Parser { return s.parser }

func (s *PropContext) CopyFrom(ctx *PropContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PropExprContext struct {
	*PropContext
}

func NewPropExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropExprContext {
	var p = new(PropExprContext)

	p.PropContext = NewEmptyPropContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropContext))

	return p
}

func (s *PropExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropExprContext) PropName() IPropNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropNameContext)
}

func (s *PropExprContext) Colon() antlr.TerminalNode {
	return s.GetToken(RippletParserColon, 0)
}

func (s *PropExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PropExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterPropExpr(s)
	}
}

func (s *PropExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitPropExpr(s)
	}
}

type ComputedPropExprContext struct {
	*PropContext
}

func NewComputedPropExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComputedPropExprContext {
	var p = new(ComputedPropExprContext)

	p.PropContext = NewEmptyPropContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropContext))

	return p
}

func (s *ComputedPropExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComputedPropExprContext) BracketOpen() antlr.TerminalNode {
	return s.GetToken(RippletParserBracketOpen, 0)
}

func (s *ComputedPropExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ComputedPropExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ComputedPropExprContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(RippletParserBracketClose, 0)
}

func (s *ComputedPropExprContext) Colon() antlr.TerminalNode {
	return s.GetToken(RippletParserColon, 0)
}

func (s *ComputedPropExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterComputedPropExpr(s)
	}
}

func (s *ComputedPropExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitComputedPropExpr(s)
	}
}

func (p *RippletParser) Prop() (localctx IPropContext) {
	localctx = NewPropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, RippletParserRULE_prop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(452)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RippletParserIntLiteral, RippletParserHexLiteral, RippletParserRealLiteral, RippletParserIdentifier, RippletParserStringOpen:
		localctx = NewPropExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(442)
			p.PropName()
		}
		{
			p.SetState(443)
			p.Match(RippletParserColon)
		}
		{
			p.SetState(444)
			p.expression(0)
		}

	case RippletParserBracketOpen:
		localctx = NewComputedPropExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(446)
			p.Match(RippletParserBracketOpen)
		}
		{
			p.SetState(447)
			p.expression(0)
		}
		{
			p.SetState(448)
			p.Match(RippletParserBracketClose)
		}
		{
			p.SetState(449)
			p.Match(RippletParserColon)
		}
		{
			p.SetState(450)
			p.expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPropNameContext is an interface to support dynamic dispatch.
type IPropNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropNameContext differentiates from other interfaces.
	IsPropNameContext()
}

type PropNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropNameContext() *PropNameContext {
	var p = new(PropNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_propName
	return p
}

func (*PropNameContext) IsPropNameContext() {}

func NewPropNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropNameContext {
	var p = new(PropNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_propName

	return p
}

func (s *PropNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PropNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RippletParserIdentifier, 0)
}

func (s *PropNameContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *PropNameContext) NumberLiteral() INumberLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberLiteralContext)
}

func (s *PropNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterPropName(s)
	}
}

func (s *PropNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitPropName(s)
	}
}

func (p *RippletParser) PropName() (localctx IPropNameContext) {
	localctx = NewPropNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, RippletParserRULE_propName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(457)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RippletParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(454)
			p.Match(RippletParserIdentifier)
		}

	case RippletParserStringOpen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(455)
			p.StringLiteral()
		}

	case RippletParserIntLiteral, RippletParserHexLiteral, RippletParserRealLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(456)
			p.NumberLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReservedWordContext is an interface to support dynamic dispatch.
type IReservedWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReservedWordContext differentiates from other interfaces.
	IsReservedWordContext()
}

type ReservedWordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedWordContext() *ReservedWordContext {
	var p = new(ReservedWordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_reservedWord
	return p
}

func (*ReservedWordContext) IsReservedWordContext() {}

func NewReservedWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedWordContext {
	var p = new(ReservedWordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_reservedWord

	return p
}

func (s *ReservedWordContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedWordContext) Keyword() IKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *ReservedWordContext) NilLiteral() INilLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INilLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INilLiteralContext)
}

func (s *ReservedWordContext) BoolLiteral() IBoolLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolLiteralContext)
}

func (s *ReservedWordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedWordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedWordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterReservedWord(s)
	}
}

func (s *ReservedWordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitReservedWord(s)
	}
}

func (p *RippletParser) ReservedWord() (localctx IReservedWordContext) {
	localctx = NewReservedWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, RippletParserRULE_reservedWord)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(462)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RippletParserIf, RippletParserThen, RippletParserElse, RippletParserMatch, RippletParserFn, RippletParserRepeat, RippletParserUntil, RippletParserIs, RippletParserIsNot, RippletParserAnd, RippletParserOr:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(459)
			p.Keyword()
		}

	case RippletParserNilLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(460)
			p.NilLiteral()
		}

	case RippletParserBooLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(461)
			p.BoolLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RippletParserRULE_keyword
	return p
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RippletParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) If() antlr.TerminalNode {
	return s.GetToken(RippletParserIf, 0)
}

func (s *KeywordContext) Then() antlr.TerminalNode {
	return s.GetToken(RippletParserThen, 0)
}

func (s *KeywordContext) Else() antlr.TerminalNode {
	return s.GetToken(RippletParserElse, 0)
}

func (s *KeywordContext) Match() antlr.TerminalNode {
	return s.GetToken(RippletParserMatch, 0)
}

func (s *KeywordContext) Fn() antlr.TerminalNode {
	return s.GetToken(RippletParserFn, 0)
}

func (s *KeywordContext) Repeat() antlr.TerminalNode {
	return s.GetToken(RippletParserRepeat, 0)
}

func (s *KeywordContext) Until() antlr.TerminalNode {
	return s.GetToken(RippletParserUntil, 0)
}

func (s *KeywordContext) Is() antlr.TerminalNode {
	return s.GetToken(RippletParserIs, 0)
}

func (s *KeywordContext) IsNot() antlr.TerminalNode {
	return s.GetToken(RippletParserIsNot, 0)
}

func (s *KeywordContext) And() antlr.TerminalNode {
	return s.GetToken(RippletParserAnd, 0)
}

func (s *KeywordContext) Or() antlr.TerminalNode {
	return s.GetToken(RippletParserOr, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RippletParserListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *RippletParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, RippletParserRULE_keyword)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(464)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RippletParserIf)|(1<<RippletParserThen)|(1<<RippletParserElse)|(1<<RippletParserMatch)|(1<<RippletParserFn)|(1<<RippletParserRepeat)|(1<<RippletParserUntil)|(1<<RippletParserIs)|(1<<RippletParserIsNot)|(1<<RippletParserAnd)|(1<<RippletParserOr))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *RippletParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RippletParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 21)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
