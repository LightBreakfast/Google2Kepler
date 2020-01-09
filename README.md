# Google2Kepler
Convert your Google Location History to a Kepler.gl Compatible CSV file

# Instuctions (Windows)

## Step 1 - Export Google Location History

* Go to https://takeout.google.com/settings/takeout

* Deselect All with exception of 'Location History'

* Select 'Next Step' (At the bottom of the page)

* On the next page, click 'Create Archive'

This will start process whereby an archive will be created containing your entire location history and then sent to you via email. For me I found it generally took around 5 minutes from request to response, but your results will vary depending on the amount of data you have. (You may also need to verify the request came from you, if this has not been done previously)

## Step 2 - Extract Archive

Once you have receieved an email from Google, download the archive via the provided link(s).

It will have the following structure

./takeout-XXXXXX/Takeout/Location History

From this directory, copy the entire "Semantic Location History" Folder, and paste it into the same Directory as G2K.exe

## Step 3 - Run Executable

Running G2K.exe will read all the JSON files in the "Semantic Location History" Folder, and output it to 'results.csv'. This should take less than a second, but times may vary depending on the quantity of data.

## Step 4 - Import Data into Kepler.gl