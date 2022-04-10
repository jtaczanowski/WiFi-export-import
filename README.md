# WiFi-export-import
--- EN ---

Program for exporting and importing WiFi netowrks profiles in Windows systems

--- PL ---

Program do importowania i eksportowania profili sieci WiFi w systemach Windows



<img src="/screenshots/WiFi-export-import_1_0_FINAL.JPG" alt="WiFi-export-import_1_0_FINAL" height="100%" width="100%">


## Usage
### Export all WiFi profiles.
- Download [WiFi-export-import_1_0_FINAL.exe](https://github.com/jtaczanowski/WiFi-export-import/raw/main/WiFi-export-import_1_0_FINAL.exe)
- Run ```WiFi-export-import_1_0_FINAL.exe``` as ⚠️administrator. (Administrator privileges are needed to export WiFi password as clear text)
- click "Eksport" button and choose directory to which you want to save exported WiFi profiles.

### Import WiFi profiles from directory.
- Run ```WiFi-export-import_1_0_FINAL.exe```.
- click "Import" button and choose directory from which you want to import WiFi profiles.
## About WiFi-export-import
--- EN ---

WiFi-export-import is a simple tool for exporting and importing WiFi networks profiles from Windows systems. The program was developed in 2012 while studing computer science and agroengineering at Poznan University of life sciences. Underthehood, for exporting and importing WiFi network profiles program is using ```netsh``` commands. Unfortunately I can't attach source code of this program because it was lost in the meantime.

--- PL ---

WiFi-export-import to prosty program pozwalający na eksportowanie i importowanie profili sieci w systemach Windows. Program powstał w 2012 roku w trakcie studiów na kierunku Informatyka i agroinżynieria (makrokierunek) na Uniwersytecie Przyrodniczym w Poznaniu. Program do eksportowania i importowania profili wykorzystuje polecenia ```netsh```. Niestety nie mogę dołączyć kodu źródłowego programu gdy ten zaginął w międzyczasie.

### System requirements
 - ✅ tested on Windows versions from 7 to 11

### Other
 - program was developed in .NET framework.
 - program uses netsh commands like ```netsh wlan export profile key=clear folder=C:\wifi``` ```wlan add profile filename=C:\wifi\WiFi-network.xml```.

