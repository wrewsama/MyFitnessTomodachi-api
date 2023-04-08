package utils

import (
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/wrewsama/MyFitnessTomodachi-api/models"
)

func GetMatches(query string, foods []models.Food) []models.Food {
	// get a list of the food names
	foodNames := make([]string, 0)
	for _, food := range foods {
		foodNames = append(foodNames, food.Name)
	}

	// use RankFind to sort by levenshtein distance
	ranks := fuzzy.RankFind(query, foodNames)

	// use the original indexes in the Rank array to get foods matching query
	res := make([]models.Food, 0)
	for _, rank := range ranks {
		currFood := foods[rank.OriginalIndex]
		res = append(res, currFood)
	}
	return res
}