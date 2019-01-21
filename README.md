# Sentiment Analysis CLI

This is a tool that can assess the sentiment of a given English sentences (currently just supports "positive" and "negative" sentiment).  It runs via the command line.  This tool is implemeted using a naive Bayes classifier that was trained on a dataset of 1000 annotated yelp reviews.  NLP for Go!

More features to come.  Some of my TODOs are adding a "mixed" sentiment rating (for sentences that have a mixture of sentiment), improving my algorithm via lemmatization (grouping like words such as "run", "ran", "running" together), chunking certain words together when nouns become adjectives (e.g. "Orange Chicken"), and potentially incorporating TF-IDF (_term frequency-inverse document frequency_, i.e. "how important is this word in the context of all of the other words in this corpus") instead of just using word frequency.

## Usage

1. Clone this repo `git clone git@github.com:dmarticus/sentiment-analyzer-cli.git`
2. Compile and build the binary by running `./build.sh`
3. Execute the CLI by running `./sentiment_analyzer`
4. ???
5. Profit

### Acknowledgements

Inspired by this fantastically descriptive post from @kcatstack [here](https://medium.com/@kcatstack/sentiment-analysis-naive-bayes-classifier-from-scratch-part-1-theory-4949115ba13)
