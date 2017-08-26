from modules import parser
from modules import requester
from tabulate import tabulate
from modules import mailfunctions
from urllib import pathname2url
from urllib import unquote
import argparse
import sys


RED   = "\033[1;31m"  
BLUE  = "\033[1;34m"
CYAN  = "\033[1;36m"
GREEN = "\033[0;32m"
RESET = "\033[0;0m"
BOLD    = "\033[;1m"
REVERSE = "\033[;7m"


linkedinUsername = "EMAIL"
linkedinPassword = "PASSWORD"


banner = """                                                                                                                  
                                                                                                                  
RRRRRRRRRRRRRRRRR                  AAA   VVVVVVVV           VVVVVVVVEEEEEEEEEEEEEEEEEEEEEENNNNNNNN        NNNNNNNN
R::::::::::::::::R                A:::A  V::::::V           V::::::VE::::::::::::::::::::EN:::::::N       N::::::N
R::::::RRRRRR:::::R              A:::::A V::::::V           V::::::VE::::::::::::::::::::EN::::::::N      N::::::N
RR:::::R     R:::::R            A:::::::AV::::::V           V::::::VEE::::::EEEEEEEEE::::EN:::::::::N     N::::::N
  R::::R     R:::::R           A:::::::::AV:::::V           V:::::V   E:::::E       EEEEEEN::::::::::N    N::::::N
  R::::R     R:::::R          A:::::A:::::AV:::::V         V:::::V    E:::::E             N:::::::::::N   N::::::N
  R::::RRRRRR:::::R          A:::::A A:::::AV:::::V       V:::::V     E::::::EEEEEEEEEE   N:::::::N::::N  N::::::N
  R:::::::::::::RR          A:::::A   A:::::AV:::::V     V:::::V      E:::::::::::::::E   N::::::N N::::N N::::::N
  R::::RRRRRR:::::R        A:::::A     A:::::AV:::::V   V:::::V       E:::::::::::::::E   N::::::N  N::::N:::::::N
  R::::R     R:::::R      A:::::AAAAAAAAA:::::AV:::::V V:::::V        E::::::EEEEEEEEEE   N::::::N   N:::::::::::N
  R::::R     R:::::R     A:::::::::::::::::::::AV:::::V:::::V         E:::::E             N::::::N    N::::::::::N
  R::::R     R:::::R    A:::::AAAAAAAAAAAAA:::::AV:::::::::V          E:::::E       EEEEEEN::::::N     N:::::::::N
RR:::::R     R:::::R   A:::::A             A:::::AV:::::::V         EE::::::EEEEEEEE:::::EN::::::N      N::::::::N
R::::::R     R:::::R  A:::::A               A:::::AV:::::V          E::::::::::::::::::::EN::::::N       N:::::::N
R::::::R     R:::::R A:::::A                 A:::::AV:::V           E::::::::::::::::::::EN::::::N        N::::::N
RRRRRRRR     RRRRRRRAAAAAAA                   AAAAAAAVVV            EEEEEEEEEEEEEEEEEEEEEENNNNNNNN         NNNNNNN

                                                                               LinkedIn Information Gathering Tool\n\n"""

# Parses the data from command line
ArgParser = argparse.ArgumentParser(description='Raven - LinkedIn Information Gathering Tool')
ArgParser.add_argument('-c','--company', help='Input the Company name. Ex: Pizzahut ', required=True)
ArgParser.add_argument('-s','--state', help='Input the State initials. Ex: uk , al , etc...', required=True)
ArgParser.add_argument('-d','--domain', help='Input the domain name. Ex: gmail.com ', required=True)
ArgParser.add_argument('-p','--pages', help='Number of google pages to navigate. Ex: 3', required=False)

# You can hardcode the credentials or use the parameters.

ArgParser.add_argument('-lu','--lusername', help='The linkedin username to use.', required=False)
ArgParser.add_argument('-lp','--lpassword', help='The linekdin password to use.', required=False)


args = vars(ArgParser.parse_args())

if args["lusername"] is not None and args["lpassword"] is not None:
	linkedinUsername = args["lusername"]
	linkedinPassword = args["lpassword"]


ParserObject = parser.Parser(args['state'])
RequesterObject = requester.Requester()

companyArg = pathname2url(args['company'])
state = args['state']

domain = args['domain']

pages_count = 1
if args["pages"] is not None:
	pages_count = args['pages']


# Prints the banner, who doesn't loves banners :P

sys.stdout.write(GREEN)
print banner
sys.stdout.write(RESET)


sys.stdout.write(CYAN)


Persons = []

# Download data from duck duck go
htmlData = RequesterObject.getLinkedinLinks(state,companyArg,pages_count)


# Parses the data from duck duck go
ParserObject.readHTMLFile(htmlData)
URLs = ParserObject.getExtractedLinks()


# Will login the requester
RequesterObject.doLogin(linkedinUsername,linkedinPassword)


sys.stdout.write(CYAN)
print "[+] Found a total of %d profiles. [+]" % len(URLs)
print "[+] Parsing all the profiles , this will take a while [+]"
for x in URLs:
	url = x.replace("https://%s." % state,"https://www.")
	print url
	response = RequesterObject.doGetLinkedin(url)
	try:
		
		name = ParserObject.extractName(response)
		position = ParserObject.extractPosition(response)
		company = ParserObject.extractCompany(response).strip()
		Person = [name,position,company]
		Persons.append(Person)

	except Exception, error:
		sys.stdout.write(RED)
		print "[-] Error : %s [-]" % error
		sys.stdout.write(CYAN)
		pass

sys.stdout.write(CYAN)

print "[+] Profiles Parsed %s [+]" % str(len(Persons))

print "[+] Cleaning Invalid Data [+]"


# CODE WRITTEN AT 3 AM

temp = []
for x in range (0,len(Persons)):
	if(str(unquote(companyArg)).lower() in str(Persons[x][2]).lower()):
		temp.append(Persons[x])
Persons = temp

# The code above will clean the profiles which are false positives

# Create the mail parser object

MailObject = mailfunctions.MailFunctions(Persons)


sys.stdout.write(RESET)

email_formats = '''

Email formats - John Doe 

#
# 1- john.doe@example.com 	-- {firstname}.{lastname}@{domain}
# 2- doe.john@example.com 	-- {lastname}.{firstname}@{domain}
# 3- john-doe@example.com 	-- {firstname}-{lastname}@{domain} 	
# 4- jdoe@example.com 		-- {firstname[0]}{lastname}@{domain}	
# 5- doe.j@example.com 		-- {lastname}{firstname[0]}@{domain} 	
# 6- d.joe@example.com 		-- {lastname[0]}{firstname}@{domain}
# 7- joe.d@example.com 		-- {firstname}{lastname[0]}@{domain}
#
'''

print email_formats
sys.stdout.write(GREEN)
format = raw_input("\n#> Enter format number: ")

sys.stdout.write(CYAN)
# Will generate emails based on pattern and will return an array
CompletedList = MailObject.generateEmails(domain,int(format))

# Will check for pwned accounts.
sys.stdout.write(RED)
MailObject.checkPwned(CompletedList)



sys.stdout.write(CYAN)
print tabulate(CompletedList, headers=['Name', 'Position', 'Company', 'E-mail'],tablefmt="fancy_grid")

# Will save the data to a CSV format for integration with phishing frameworks
sys.stdout.write(GREEN)
MailObject.saveOutput(CompletedList,domain)

RequesterObject.kill()

print "\n @0x09AL - https://twitter.com/0x09AL\n"