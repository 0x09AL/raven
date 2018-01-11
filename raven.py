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

# add the R before banner text
banner = r"""
__________                                     _______       ________
\______   \_____ ___  __ ____   ____           \   _  \      \_____  \
 |       _/\__  \\  \/ // __ \ /    \   ______ /  /_\  \      /  ____/
 |    |   \ / __ \\   /\  ___/|   |  \ /_____/ \  \_/   \    /       \
 |____|_  /(____  /\_/  \___  >___|  /          \_____  / /\ \_______ \
        \/      \/          \/     \/                 \/  \/         \/
                                                                               LinkedIn Information Gathering Tool - by @0x09AL\n\n"""



# Parses the data from command line
ArgParser = argparse.ArgumentParser(description='Raven - LinkedIn Information Gathering Tool')
ArgParser.add_argument('-c','--company', help='Input the Company name. Ex: Pizzahut ', required=True)
ArgParser.add_argument('-s','--state', help='Input the State initials. Ex: uk , al , etc...', required=True)
ArgParser.add_argument('-d','--domain', help='Input the domain name. Ex: gmail.com ', required=False)
ArgParser.add_argument('-p','--pages', help='Number of google pages to navigate. Ex: 3', required=False)
ArgParser.add_argument('-f','--format', help='Specify format type. Ex: 1,2 or ALL', required=False)
ArgParser.add_argument('-v','--verify', help='Verify e-mails by using OWA. Ex: https://mail.example.com/', required=False)
ArgParser.add_argument('-l','--list', help='List formats', required=False, action='store_true')
ArgParser.add_argument('-chp','--check-pwned', help='Checks if the email can be found in a public databreach', required=False, action='store_true')

# You can hardcode the credentials or use the parameters.
ArgParser.add_argument('-lu','--lusername', help='The linkedin username to use.', required=False)
ArgParser.add_argument('-lp','--lpassword', help='The linekdin password to use.', required=False)


args = vars(ArgParser.parse_args())

owa_url = args['verify']
email_format = args['format']


email_formats = '''
Email formats - John Doe

# 1- john.doe@example.com 	-- {firstname}.{lastname}@{domain}
# 2- doe.john@example.com 	-- {lastname}.{firstname}@{domain}
# 3- john-doe@example.com 	-- {firstname}-{lastname}@{domain}
# 4- jdoe@example.com 		-- {firstname[0]}{lastname}@{domain}
# 5- doe.j@example.com 		-- {lastname}{firstname[0]}@{domain}
# 6- d.joe@example.com 		-- {lastname[0]}{firstname}@{domain}
# 7- joe.d@example.com 		-- {firstname}{lastname[0]}@{domain}
'''

if(args["list"]):
        print email_formats
        exit(0)

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
sys.stdout.write(CYAN)
print banner

Persons = []

# Download data from duck duck go
htmlData = RequesterObject.getLinkedinLinks(state,companyArg,pages_count)


# Parses the data from duck duck go
ParserObject.readHTMLFile(htmlData)
URLs = ParserObject.getExtractedLinks()


# Will login the requester
if(not RequesterObject.doLogin(linkedinUsername,linkedinPassword)):
    RequesterObject.kill()
    exit(0)


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

sys.stdout.write(CYAN)
# Will generate emails based on pattern and will return an array
CompletedList = MailObject.generateEmails(domain,int(email_format))


# Will check for pwned accounts.
if(args["check_pwned"]):


    MailObject.checkPwned(CompletedList)



sys.stdout.write(CYAN)
print tabulate(CompletedList, headers=['Name', 'Position', 'Company', 'E-mail'],tablefmt="fancy_grid")

# Will save the data to a CSV format for integration with phishing frameworks

MailObject.saveOutput(CompletedList,domain)
RequesterObject.kill()

print "\n @0x09AL - https://twitter.com/0x09AL\n"
