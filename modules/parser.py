from bs4 import BeautifulSoup
import requests

class Parser(object):

	company = ""
	linkedInUrl = ""
	htmlData = ""
	linkedinURLS = []


	def __init__(self,country):

		self.linkedInUrl = "https://%s.linkedin.com/in/" % country
		#self.company = company


	def readHTMLFile(self,htmlData):
		# This will initialize the html data
		self.htmlData = htmlData

	def getCompany():
		# Returns the company name
		return self.company

	def getExtractedLinks(self):
		# Returns Extracted Links
		print "[+] Parsing Data from html file [+]"

		if(self.htmlData == ""):
			print "[-] There is no html data exiting [-]"
			exit()
		soupParser = BeautifulSoup(self.htmlData, 'html.parser')

		for link in soupParser.find_all('a'):
			
			temp = str(link.get('href'))
			
			if(temp.startswith(self.linkedInUrl)):
				if(temp not in self.linkedinURLS):
					self.linkedinURLS.append(temp)
					
		return self.linkedinURLS
		# Return linkedinURLS array



	def getEmployeeInformation(self):
		# Will return the dictionary that contains employee data
		return


# The response parameter is the data that every visited link will response basically the html page of the persons linkedin

	def extractName(self,response):
		# Will return the name and the surname from the response
		soupParser = BeautifulSoup(response, 'html.parser')
		name = soupParser.findAll("h1", class_="pv-top-card-section__name")[0].string

		return name

	def extractPosition(self,response):
		soupParser = BeautifulSoup(response, 'html.parser')
		position = soupParser.findAll("h2", class_="pv-top-card-section__headline")[0].string
		# To avoid big position names that will break the table format.
		if(len(position)>40):
			position = position[0:40]
		
		return position

	def extractCompany(self,response):
		# Will return the company from the response
		soupParser = BeautifulSoup(response, 'html.parser')
		company = soupParser.findAll("h3", class_="pv-top-card-section__company")[0].string
		return company
		


	def extractPhone(self,response):
		# Will return the phone if found
		# To be implemented
		return





