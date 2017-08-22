import time
import requests
import json

class MailFunctions(object):


	Persons = []

	def __init__(self,persons):

		self.Persons = persons
		print "[+] Mail functions initialized [+]"

	def generateEmails(self,domain,format):
		# Email formats
		# 1-{firstname}.{lastname}@{domain}
		# 2-{lastname}.{firstname}@{domain}
		# 3-{firstname}-{lastname}@{domain}
		# 4-{firstname[0]}{lastname}@{domain}
		# 5-{lastname}{firstname[0]}@{domain}
		# 6-{lastname[0]}{firstname}@{domain}
		# 7-{firstname}{lastname[0]}@{domain}
		for person in self.Persons:
			try:
				firstname = person[0].split(" ")[0]
				lastname = person[0].split(" ")[1]
			except Exception, error:
				print "[-] Error: %s [-]" % error
			
			if(format==1):
				email = "%s.%s@%s" % (firstname,lastname,domain)
			elif(format==2):
				email = "%s.%s@%s" % (lastname,firstname, domain)
			elif(format==3):
				email = "%s-%s@%s" % (firstname,lastname, domain)
			elif(format==4):
				email = "%s%s@%s" % (firstname[0],lastname,domain)
			elif(format==5):
				email = "%s.%s@%s" % (lastname,firstname[0],domain)
			elif(format==6):
				email = "%s.%s@%s" % (firstname,lastname[0],domain)
			elif(format==7):
				email = "%s.%s@%s" % (lastname[0],firstname,domain)
			else:
				print "[-] Invalid Option [-]" # Normally we should never come here
			#	return
			person.append(email.lower())
		return self.Persons

	def saveOutput(self,emailList):
		
		filename = "%s%s%s" % ("output/",str(int(time.time())),'.csv')
		print "[+] Saving output to %s " % filename
		output = open(filename , "w")
		head = "First Name,Last Name,Position,Email\n"
		output.write(head)

		for person in range (0,len(emailList)):
			line = "%s,%s,%s,%s\n" % (emailList[person][0].split(" ")[0],emailList[person][0].split(" ")[1],emailList[person][1],emailList[person][3])
			output.write(line)

		output.close()

	def checkPwned(self,emailList):

		for email in emailList:
			url = "https://haveibeenpwned.com/api/v2/breachedaccount/%s?truncateResponse=true" % email[3]
			time.sleep(1) # Sleep to avoid many requests error from the web server

			r = requests.get(url,headers={'User-Agent':'Mozilla/5.0 (Windows NT 10.0; WOW64; rv:40.0) Gecko/20100101 Firefox/40.0'})
			if(r.status_code == 200):
				try:
					pwnedSites = json.loads(r.text)
					for site in pwnedSites:
						print "[+] %s pwned at %s breach [+]" % (email[3],site["Name"])
				except Exception:
					pass
				
			