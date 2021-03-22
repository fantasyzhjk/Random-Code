require 'selenium-webdriver' 
driver = Selenium::WebDriver.for :chrome
driver.get "https://www.arealme.cn/apm-actions-per-minute-test/cn/" 
sleep 1

driver.find_element(:id, 'start').click
50.times{driver.find_element(:class, 'atm-latest').click}

sleep

