# -*- coding: utf-8 -*-
import pyautogui 
import subprocess 
from time import sleep


class XY:

    def __init__(self, x = 0, y = 0):
        self.x = x
        self.y = y

def main():
    pyautogui.FAILSAFE = False
    dispsize = XY(pyautogui.size()[0], pyautogui.size()[1])
    center = XY(dispsize.x // 2, dispsize.y // 2)

    # 絶対座標指定
    pyautogui.moveTo(center.x, center.y)
    sleep(1)

    # 相対座標指定
    pyautogui.moveRel(-center.x, -center.y, 1)

    pyautogui.click(button="right")
    sleep(1)

    pyautogui.hotkey('esc')
    sleep(1)

    pyautogui.alert("test")

    ssname = "tmp.png"
    subprocess.call(["ls"])
    sleep(1)
    pyautogui.screenshot(ssname, region=(100, 100, 50, 50))
    subprocess.call(["ls"])
    sleep(1)

    pyautogui.moveTo(pyautogui.locateCenterOnScreen(ssname))
    subprocess.call(["rm", "-rf", ssname])


if __name__ == '__main__':
    main()
