# 🛍️ OLX.ua Scraper & Statistics Generator 🛍️

Welcome to the **OLX.ua Scraper & Statistics Generator** project! 🎉 <br> 
This tool allows you to scrape data from [OLX.ua](https://www.olx.ua/), collect valuable information about listings, and generate insightful statistics. 📊

## 🚀 Features

- **Scrape Categories**: Automatically fetches all categories from OLX.ua. 📂
- **Collect Listings**: Gathers detailed information about each listing, including ID, title, price, location, and posting date. 📝
- **Track Sold Items**: Monitors listings over time to determine when items are sold. ⏳
- **Generate Statistics**: Calculates average selling times per category. 📈
- **Data Persistence**: Saves data to a CSV file for further analysis. 💾

## 🛠️ Installation

### Prerequisites

- [Go](https://golang.org/dl/) installed (version 1.16 or higher). 🐹
- [Git](https://git-scm.com/downloads) installed. 🌳

### Clone the Repository

```bash
git clone https://github.com/yourusername/olx-scraper.git
cd olx-scraper
```

### Install Dependencies
This project uses the colly library for web scraping.
```bash 
 go get -u github.com/gocolly/colly/...
```
## 🎯 Usage
Run the scraper using the following command:
```bash
go run main.go
```
### The scraper will:

1) Visit the OLX.ua homepage and collect all category URLs. 🌐
2) Iterate through each category and scrape listings. 🔍
3) Extract relevant data from each listing. 🧮
4) Save the data to items.csv. 📂
5) Generate statistics based on the collected data. 📊
## 📄 Data Fields
### Each listing collected will have the following fields:
* OlxID: Unique identifier for the listing. 🆔
* Category: The category under which the listing is posted. 🗂️
* Title: The title of the listing. 📝
* Price: The price of the item. 💰
* Location: The location of the item. 📍
* DatePosted: The date and time when the listing was posted. 📅
* DateSold: The date and time when the item was sold (if applicable). ⌛
## 🧩 Project Structure

* ```main.go```: The main program file containing all the logic. ⚙️
* ```items.csv```: The CSV file where all listing data is stored. 💾
## 🔍 How It Works
1) Category Extraction: The scraper uses a CSS selector to find all category links on the OLX.ua homepage.
2) Listing Extraction: For each category, it visits the category page and uses a CSS selector to find all listings.
3) Data Parsing:
* Title & Price: Extracts the title and price using child selectors.
* Location & Date: Parses the location and date information from the listing details.
* Date Parsing: Handles different date formats (e.g., "Today", "Yesterday", "15 October at 15:04"). 📆
4) Data Storage: All extracted data is stored in a map and then written to items.csv.
5) Statistics Generation: Calculates the average time it takes for items to be sold in each categor
## 📊 Sample Statistics Output
```
Category: Electronics
Items Sold: 50
Average Selling Time: 72.00 hours

Category: Fashion
Items Sold: 30
Average Selling Time: 48.00 hours
```
## ⚠️ Important Notes
* Respect the Website's Terms of Service: Ensure that scraping OLX.ua complies with their terms of service. ❗
* Be Polite: The scraper includes delays to avoid overwhelming the server. Adjust RandomDelay as needed. ⏱️
* Data Accuracy: The accuracy of the scraped data depends on the consistency of the website's structure. 🧐

## 🤝 Contributing
Contributions are welcome! 
If you have suggestions or improvements, feel free to open an issue or submit a pull request. 💡

## 🛡️ Disclaimer
This project is intended for educational purposes. The author is not responsible for any misuse of this tool. Use it responsibly. 🛑