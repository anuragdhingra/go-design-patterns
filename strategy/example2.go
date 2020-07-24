package main

import (
	"bufio"
	"fmt"
	"os"
)

// =========================== ANALYZER =================================
type Analyzer interface {
	Analyze(string) string
}

type kagomeAnalyzer struct{}

func (kagome kagomeAnalyzer) Analyze(query string) string {
	return fmt.Sprintf("You are using kagome analyzer for %s", query)
}

type icuAnalyzer struct{}

func (icu icuAnalyzer) Analyze(query string) string {
	return fmt.Sprintf("You are using icu analyzer for %s", query)
}

// ======================================================================

// =========================== SEARCHER =================================
type Searcher interface {
	Search(string) string
}

type esSearch struct {
	address string
}

func (elasticsearch esSearch) Search(query string) (results string) {
	return fmt.Sprintf("You are using elasticsearch at endpoint: %s for query: %s", elasticsearch.address, query)
}

type solrSearch struct {
	address string
}

func (solr solrSearch) Search(query string) (results string) {
	return fmt.Sprintf("You are using Solr search at endpoint: %s for query: %s", solr.address, query)
}

// ======================================================================

// =========================== METADATA FETCHER =================================
type MetadataFetcher interface {
	Metadata(string) string
}

type noFetcher struct{}

func (nf noFetcher) Metadata(query string) string {
	return fmt.Sprintf("sorry we dont have any metadata for %s", query)
}

type spacyFetcher struct{}

func (sf spacyFetcher) Metadata(query string) string {
	return fmt.Sprintf("you are using spacy fetcher for %s", query)
}

// ======================================================================

// ======================================================================
type searchEngine struct {
	analyzer        Analyzer
	searcher        Searcher
	metadataFetcher MetadataFetcher
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter version to use [v1, v2, v3, v4]: ")
	version, _ := reader.ReadString('\n')
	fmt.Printf("Version choosed: %s", version)
	fmt.Print("Enter search query: ")
	query, _ := reader.ReadString('\n')
	fmt.Printf("query is: %s", query)

	v1 := searchEngine{
		analyzer: kagomeAnalyzer{},
		searcher: esSearch{
			address: "localhost:9000",
		},
		metadataFetcher: noFetcher{},
	}

	v2 := searchEngine{
		analyzer: kagomeAnalyzer{},
		searcher: solrSearch{
			address: "localhost:8983",
		},
		metadataFetcher: noFetcher{},
	}

	v3 := searchEngine{
		analyzer: icuAnalyzer{},
		searcher: esSearch{
			address: "localhost:9000",
		},
		metadataFetcher: spacyFetcher{},
	}

	v4 := searchEngine{
		analyzer:
		icuAnalyzer{},
		searcher: solrSearch{
			address: "localhost:8983",
		},
		metadataFetcher: spacyFetcher{},
	}

	fmt.Println(version)

	if version == "v1\n" {
		fmt.Printf("Query is: %s", query)
		fmt.Printf("Analyzed: %s", v1.analyzer.Analyze(query))
		fmt.Printf("Search results: %s", v1.searcher.Search(query))
		fmt.Printf("Query Metadata: %s", v1.metadataFetcher.Metadata(query))
	} else if version == "v2\n" {
		fmt.Printf("Query is: %s", query)
		fmt.Printf("Analyzed: %s", v2.analyzer.Analyze(query))
		fmt.Printf("Search results: %s", v2.searcher.Search(query))
		fmt.Printf("Query Metadata: %s", v2.metadataFetcher.Metadata(query))
	} else if version == "v3\n" {
		fmt.Printf("Query is: %s", query)
		fmt.Printf("Analyzed: %s", v3.analyzer.Analyze(query))
		fmt.Printf("Search results: %s", v3.searcher.Search(query))
		fmt.Printf("Query Metadata: %s", v3.metadataFetcher.Metadata(query))
	} else if version == "v4\n" {
		fmt.Printf("Query is: %s", query)
		fmt.Printf("Analyzed: %s", v4.analyzer.Analyze(query))
		fmt.Printf("Search results: %s", v4.searcher.Search(query))
		fmt.Printf("Query Metadata: %s", v4.metadataFetcher.Metadata(query))
	} else {
		fmt.Println("Please specify correct version")
	}
}

// Youtube link: https://www.youtube.com/watch?v=v9ejT8FO-7I

