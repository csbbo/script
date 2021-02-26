from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By
import time

# http://chromedriver.chromium.org/downloads

if __name__ == '__main__':
    browser = webdriver.Chrome()
    browser.implicitly_wait(10)
    browser.get("https://www.aliyun.com")
    logined = input('完成登录操作: ')
    try:
        browser.get("https://www.aliyun.com/?spm=a2c44.11131956.0.0.a88d54558aBxjE&accounttraceid=9153756880924eadafb932273a5a2806nkbn")
        browser.find_element_by_link_text('控制台').click()

        browser.switch_to.window(browser.window_handles[-1])
        browser.find_element_by_link_text('工单').click()

        search_box = browser.find_elements(By.XPATH, '/html/body/div[1]/div/div[1]/div[2]/div[2]/div/div/div[1]/div/span[2]/input')[0]
        search_box.send_keys('cdn加速')
        search_box.send_keys(Keys.RETURN)

        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        pass
    except Exception as e:
        print(str(e))
    finally:
        browser.quit()
