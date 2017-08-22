# raven
raven is a Linkedin information gathering tool that  can be used by pentesters to gather information about an organization employees using Linkedin.

Please do not use this program to do stupid things. The author does not keep any responsability of what damage has been done by this program.

USE IT AT YOUR OWN RISK.

# Installation

* Run setup.sh as root
* Adjust Linkedin credentials in raven.py

# Documentation

Usage of this is application is pretty simple.
It requires two parameters. The first one is the company name and the second one is the country initials.

For example , if the company that you want to search is Evil Corp and the state is Albania the parameters would be:

python raven.py -c 'Evil Corp' -s al

# Features

* Automatically check found emails in haveibeenpwned.com
* Output in CSV format (For using with GoPhish)

# Screenshots

Screenshot - 1 - 

![ScreenShot](https://raw.githubusercontent.com/0x09AL/raven/master/screenshots/screenshot-01.png)


Screenshot - 2

![ScreenShot](https://raw.githubusercontent.com/0x09AL/raven/master/screenshots/screenshot-02.png)


Screenshot - 3

![ScreenShot](https://raw.githubusercontent.com/0x09AL/raven/master/screenshots/screenshot-03.png)


Screenshot - 4

![ScreenShot](https://raw.githubusercontent.com/0x09AL/raven/master/screenshots/screenshot-04.png)


Screenshot - 5

![ScreenShot](https://raw.githubusercontent.com/0x09AL/raven/master/screenshots/screenshot-05.png)
