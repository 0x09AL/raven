from selenium import webdriver
from pyvirtualdisplay import Display
from time import sleep
import sys


RED   = "\033[1;31m"
BLUE  = "\033[1;34m"
CYAN  = "\033[1;36m"
GREEN = "\033[0;32m"
RESET = "\033[0;0m"
BOLD    = "\033[;1m"
REVERSE = "\033[;7m"

class Requester(object):

	timeout = 10

	def __init__(self):
		display = Display(visible=0, size=(1600, 1024))
		display.start()
		self.driver = webdriver.Firefox()
		self.driver.delete_all_cookies()



	def doLogin(self,username,password):

		self.driver.get("https://www.linkedin.com/uas/login")
		self.driver.execute_script('localStorage.clear();')
		# Fixed UTF-8 issue in title.

		if(str(self.driver.title).encode('ascii','replace').startswith("Sign In")):
			print "[+] Login Page loaded successfully [+]"
			lnkUsername = self.driver.find_element_by_id("session_key-login")
			lnkUsername.send_keys(username)
			lnkPassword = self.driver.find_element_by_id("session_password-login")
			lnkPassword.send_keys(password)
			self.driver.find_element_by_id("btn-primary").click()
			sleep(5)
			if(str(self.driver.title) == "LinkedIn"):
				sys.stdout.write(CYAN)
				print "[+] Login Success [+]"
				sys.stdout.write(RESET)
				return True
			else:
				sys.stdout.write(RED)
				print "[-] Login Failed [-]"
				sys.stdout.write(RESET)
				return False



	def doGetLinkedin(self,url):
		self.driver.get(url)
		sleep(3)
		# Fix this with a better error Handling
		return self.driver.page_source.encode('ascii','replace')

	def getLinkedinLinks(self,state,company,pages_count=1):
		print "[+] Getting profiles from Google [+]"
		dork = "site:%s.linkedin.com Current: %s" % (state , company)


		self.driver.get("https://www.google.com/search?q=%s&t=h_&ia=web" % dork)
		data = self.driver.page_source.encode('ascii','replace')
		if(pages_count > 1):
			for i in range(1,int(pages_count)):
				start_at = 10 * i
				print "[+] Checking page %d on Google [+]" % i
				self.driver.get("https://www.google.com/search?q=%s&t=h_&ia=web&start=%d" % (dork,start_at))
				sleep(1)
				data += self.driver.page_source.encode('ascii','replace')
		return data

	def kill(self):
		self.driver.quit()
