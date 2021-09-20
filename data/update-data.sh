#!/bin/bash

curl https://www.oenb.at/docroot/downloads_observ/sepa-zv-vz_gesamt.csv | iconv -f ISO-8859-1 -t UTF-8 > at.csv
curl -L https://www.bundesbank.de/resource/blob/602632/de8bc9ff83a50360ade75055734724aa/mL/blz-aktuell-txt-data.txt | iconv -f ISO-8859-1 -t UTF-8 > bundesbank.txt
curl https://api.six-group.com/api/epcd/bankmaster/v2/public/downloads/bcbankenstamm_d.xls -o ch.xls && libreoffice --convert-to xlsx ch.xls --headless
 