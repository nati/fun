import copy
import math
import re
import subprocess
import sys
import time

ret = subprocess.check_output(["resize"])
m = re.match("COLUMNS=(\d+);\nLINES=(\d+);", ret)
WIDTH = int(m.group(1))
HEIGHT = int(m.group(2))

SCALE = 7

X = 0
Y = 1
Z = 2

POINTS = [
    [-1, -1, 1],
    [-1, 1, 1],
    [1, 1, 1],
    [1, -1, 1],
    [-1, -1, -1],
    [-1, 1, -1],
    [1, 1, -1],
    [1, -1, -1]
]

LINES = [
    [0, 1],
    [1, 2],
    [2, 3],
    [0, 3],
    [4, 5],
    [5, 6],
    [6, 7],
    [7, 4],
    [0, 4],
    [1, 5],
    [2, 6],
    [3, 7],
]

POINTS2 = [
    [-1, -1, 0],
    [-1, 1, 0],
    [1, 1, 0],
    [1, -1, 0],
    [0, 0, 3],
]

LINES2 = [
    [0, 1],
    [1, 2],
    [2, 3],
    [3, 0],
    [0, 4],
    [1, 4],
    [2, 4],
    [3, 4]
]


class Campas(object):
    def draw_line(self, p1, p2):
        steep = abs(p2[Y] - p1[Y]) > abs(p2[X] - p1[X])
        if steep:
            p1[X], p1[Y] = p1[Y], p1[X]
            p2[X], p2[Y] = p2[Y], p2[X]
        if p1[X] > p2[X]:
            p1[X], p2[X] = p2[X], p1[X]
            p1[Y], p2[Y] = p2[Y], p1[Y]
        dx = p2[X] - p1[X]
        dy = abs(p2[Y] - p1[Y])
        error = dx / 2.0
        y = p1[Y]
        if p1[Y] < p2[Y]:
            ystep = 1
        else:
            ystep = -1

        for x in range(p1[X], p2[X]):
            if steep:
                self.draw_point([y, x])
            else:
                self.draw_point([x, y])
            error = error - dy
            if error < 0:
                y = y + ystep
                error = error + dx

    def draw_point(self, p, char="#"):
        if p[X] >= WIDTH or 0 > p[X]:
            return
        if p[Y] >= HEIGHT or 0 > p[Y]:
            return

        sys.stdout.write("\033[%i;%iH%s" % (p[Y], p[X], char))

    def clear_screen(self):
        sys.stdout.write("\033[2J")

    def flush(self):
        sys.stdout.flush()


class Poly(object):
    points = []
    lines = []

    def __init__(self, points, lines, campas):
        self.points = copy.deepcopy(points)
        self.lines = copy.deepcopy(lines)
        self.campas = campas
        self.base_point = [0, 0, 1]

    def mult(self, transform):
        self.points = [self.mult_m_p(transform, p) for p in self.points]

    def move(self, axis, distance):
        self.base_point[axis] = distance

    def mult_m_p(self, m, p):
        x, y, z = p
        r1 = sum([m[0][0] * x, m[0][1] * y, m[0][2] * z])
        r2 = sum([m[1][0] * x, m[1][1] * y, m[1][2] * z])
        r3 = sum([m[2][0] * x, m[2][1] * y, m[2][2] * z])
        return [r1, r2, r3]

    def projection(self, p):
        cx, cy = WIDTH / 2, HEIGHT / 2
        x = (p[X] + self.base_point[X]) * SCALE / self.base_point[Z] + cx
        y = (p[Y] + self.base_point[Y]) * SCALE / self.base_point[Z] + cy
        return [int(x), int(y)]

    def draw(self):
        if self.base_point[Z] <= 0:
            return

        for point in self.points:
            self.campas.draw_point(self.projection(point))

        for line in self.lines:
            self.campas.draw_line(self.projection(self.points[line[0]]),
                                  self.projection(self.points[line[1]]))


def matrix_rotate_x(a):
    return [[1, 0, 0],
            [0, math.cos(a), -math.sin(a)],
            [0, math.sin(a), math.cos(a)]]


def matrix_rotate_y(a):
    return [[math.cos(a), 0, math.sin(a)],
            [0, 1, 0],
            [-math.sin(a), 0, math.cos(a)]]

campas = Campas()
campas.clear_screen()
cube = Poly(POINTS, LINES, campas)
cube2 = Poly(POINTS2, LINES2, campas)
cube3 = Poly(POINTS, LINES, campas)

i = math.pi / 100.0
j = 0
mx = matrix_rotate_x(i * 1)
my = matrix_rotate_y(i * 5)

while True:
    campas.clear_screen()
    cube.mult(mx)
    cube.mult(my)
    cube3.mult(mx)
    cube3.mult(my)
    cube.move(Z, math.sin(j) + 1.5)
    cube.move(X, 10 * math.cos(j))
    cube3.move(Z, math.sin(j + math.pi / 2) + 1.5)
    cube3.move(Y, 3 * math.cos(j + math.pi / 2))
    j += math.pi / 50.0
    cube2.mult(mx)
    cube2.mult(my)
    cube2.move(Z, 1.5)
    cube.draw()
    cube2.draw()
    cube3.draw()
    campas.flush()
    time.sleep(0.1)
