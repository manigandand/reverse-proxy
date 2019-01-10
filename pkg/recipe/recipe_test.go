package recipe_test

import (
	. "reverse-proxy/pkg/recipe"
	"sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Recipe", func() {
	Context("Test MustEnv", func() {
		var RecipesCon RecipesSort
		TestIngredients := []Ingredient{
			Ingredient{
				Name:      "Chicken Breasts",
				ImageLink: "https://d3hvwccx09j84u.cloudfront.net/200,200/image/554a2efafd2cb9ce488b4567.png",
			},
			Ingredient{
				Name:      "Southwest Blend Spice",
				ImageLink: "https://d3hvwccx09j84u.cloudfront.net/200,200/image/55a55c56fd2cb9f27c8b4567.png",
			},
			Ingredient{
				Name:      "Bell Pepper",
				ImageLink: "https://d3hvwccx09j84u.cloudfront.net/200,200/ingredients/596631bfc9fd08350247b2f2-aa0ce52c.png",
			},
			Ingredient{
				Name:      "Scallions",
				ImageLink: "https://d3hvwccx09j84u.cloudfront.net/200,200/image/554a301f4dab71626c8b4569.png",
			},
			Ingredient{
				Name:      "Enchilada Sauce",
				ImageLink: "https://d3hvwccx09j84u.cloudfront.net/200,200/ingredient/5aa05b45ae08b57b4b5b5471-3bc147cc.png",
			},
			Ingredient{
				Name:      "Flour Tortillas",
				ImageLink: "https://d3hvwccx09j84u.cloudfront.net/200,200/ingredient/596f6137d56afa3be43b24e2-166cf776.png",
			},
			Ingredient{
				Name:      "Monterey Jack Cheese",
				ImageLink: "https://d3hvwccx09j84u.cloudfront.net/200,200/image/55c109e74dab7112098b4567.png",
			},
			Ingredient{
				Name:      "Lime",
				ImageLink: "https://d3hvwccx09j84u.cloudfront.net/200,200/ingredient/554a3c9efd2cb9ba4f8b456c-f32287bd.png",
			},
			Ingredient{
				Name:      "Sour Cream",
				ImageLink: "https://d3hvwccx09j84u.cloudfront.net/200,200/image/5550e1064dab71893e8b4569.png",
			},
			Ingredient{
				Name:      "Olive Oil",
				ImageLink: "https://d3hvwccx09j84u.cloudfront.net/200,200/image/5566cdf2f8b25e0d298b4568.png",
			},
			Ingredient{
				Name:      "Vegetable Oil",
				ImageLink: "https://d3hvwccx09j84u.cloudfront.net/200,200/ingredients/5566d4f94dab715a078b4568-7c93a003.png",
			},
			Ingredient{
				Name:      "Salt",
				ImageLink: "https://d3hvwccx09j84u.cloudfront.net/200,200/image/5566ceb7fd2cb95f7f8b4567.png",
			},
			Ingredient{
				Name:      "Pepper",
				ImageLink: "https://d3hvwccx09j84u.cloudfront.net/200,200/image/5566dc00f8b25e5b298b4568.png",
			},
		}
		BeforeEach(func() {
		})
		AfterEach(func() {
			RecipesCon = nil
		})

		It("Should sort the recipes by preparation time in asc", func() {
			recipes := []*Recipe{
				&Recipe{
					ID:          14,
					Name:        "Cheesy Chicken Enchilada Bake",
					Headline:    "with Bell Peppers and Monterey Jack Cheese",
					Description: "It’s time to take your Taco Tuesday rotation to the next level with a tortilla casserole bake. This version is essentially an oversized sandwich of melty, saucy, and meaty flavors, featuring a layer of tortillas topped with chicken and bell peppers that then get crowned with another layer of tortillas, plus a tomato enchilada sauce and a sprinkling of cheese. When it comes out of the oven, it’s bubbling, piping hot, and steamy—how could it not impress?",
					Difficulty:  1,
					PrepTime:    "PT45M",
					ImageLink:   "https://d3hvwccx09j84u.cloudfront.net/0,0/image/cheesy-chicken-enchilada-bake-bae7dd47.jpg",
					Ingredients: TestIngredients,
				},
				&Recipe{
					ID:          1,
					Name:        "Parmesan-Crusted Pork Tenderloin",
					Headline:    "with Potato Wedges and Apple Walnut Salad",
					Description: "Parm’s the charm with this next-level pork recipe. The cheese is mixed with panko breadcrumbs to create a crust that coats the tenderloin like a glorious golden-brown crown. That way, you get meltiness, juiciness, and crunch in every bite. But this recipe isn’t just about the meat: there’s also roasted rosemary potatoes and a crisp apple walnut salad to round things out.",
					Difficulty:  1,
					PrepTime:    "PT30M",
					ImageLink:   "https://d3hvwccx09j84u.cloudfront.net/0,0/image/parmesan-crusted-pork-tenderloin-66608000.jpg",
					Ingredients: TestIngredients,
				},
				&Recipe{
					ID:          3,
					Name:        "Tex-Mex Tilapia",
					Headline:    "with Cilantro Lime Couscous and Green Beans",
					Description: "Let’s take tilapia to the next level and turn it into a Tex-Mex-style triumph on your plate. The firm-fleshed fish is given a dusting of panko breadcrumbs and our Southwest spice blend, which ensures that it has satisfying crunch and zesty flavor in every bite. The couscous, green beans, and lime crema on the side come together in a flash, meaning you’ll have this masterpiece of a meal on the table in a matter of minutes.",
					Difficulty:  1,
					PrepTime:    "PT20M",
					ImageLink:   "https://d3hvwccx09j84u.cloudfront.net/0,0/image/5a958c0d30006c33ca2850f2-c352c2d5.jpg",
					Ingredients: TestIngredients,
				},
				&Recipe{
					ID:          23,
					Name:        "Meatloaf à La Mom",
					Headline:    "with Roasted Root Veggies and Thyme Gravy",
					Description: "It’s commonly agreed that no one makes meatloaf quite like Mom. But our chefs will settle for second best by taking a page from her book and passing on this recipe that stays about as true to tradition as can be. These ground beef mini loaves are brushed with a ketchup glaze, served with tender roasted carrots and potatoes, and drizzled with an herby thyme gravy. If you’re a mom (or pop) who wants to claim this recipe as your own, don’t worry—we’ll stay mum.",
					Difficulty:  1,
					PrepTime:    "PT35M",
					ImageLink:   "https://d3hvwccx09j84u.cloudfront.net/0,0/image/5a8f0db7ae08b52cf0617622-a8d742cd.jpg",
					Ingredients: TestIngredients,
				},
			}
			RecipesCon = append(RecipesCon, recipes...)
			sort.Sort(RecipesCon)
			// for _, r := range RecipesCon {
			// 	fmt.Println(r.ID, " ==> ", r.PrepTime)
			// }
			Expect(RecipesCon[0].ID).To(Equal(3))
			Expect(RecipesCon[0].PrepTime).To(Equal("PT20M"))
			Expect(RecipesCon[1].ID).To(Equal(1))
			Expect(RecipesCon[1].PrepTime).To(Equal("PT30M"))
			Expect(RecipesCon[2].ID).To(Equal(23))
			Expect(RecipesCon[2].PrepTime).To(Equal("PT35M"))
			Expect(RecipesCon[3].ID).To(Equal(14))
			Expect(RecipesCon[3].PrepTime).To(Equal("PT45M"))
		})
		It("Should sort the recipes by preparation time in asc", func() {
			recipes := []*Recipe{
				&Recipe{
					ID:          14,
					Name:        "Cheesy Chicken Enchilada Bake",
					Headline:    "with Bell Peppers and Monterey Jack Cheese",
					Description: "It’s time to take your Taco Tuesday rotation to the next level with a tortilla casserole bake. This version is essentially an oversized sandwich of melty, saucy, and meaty flavors, featuring a layer of tortillas topped with chicken and bell peppers that then get crowned with another layer of tortillas, plus a tomato enchilada sauce and a sprinkling of cheese. When it comes out of the oven, it’s bubbling, piping hot, and steamy—how could it not impress?",
					Difficulty:  1,
					PrepTime:    "PT45M",
					ImageLink:   "https://d3hvwccx09j84u.cloudfront.net/0,0/image/cheesy-chicken-enchilada-bake-bae7dd47.jpg",
					Ingredients: TestIngredients,
				},
				&Recipe{
					ID:          1,
					Name:        "Parmesan-Crusted Pork Tenderloin",
					Headline:    "with Potato Wedges and Apple Walnut Salad",
					Description: "Parm’s the charm with this next-level pork recipe. The cheese is mixed with panko breadcrumbs to create a crust that coats the tenderloin like a glorious golden-brown crown. That way, you get meltiness, juiciness, and crunch in every bite. But this recipe isn’t just about the meat: there’s also roasted rosemary potatoes and a crisp apple walnut salad to round things out.",
					Difficulty:  1,
					PrepTime:    "PT35M",
					ImageLink:   "https://d3hvwccx09j84u.cloudfront.net/0,0/image/parmesan-crusted-pork-tenderloin-66608000.jpg",
					Ingredients: TestIngredients,
				},
				&Recipe{
					ID:          3,
					Name:        "Tex-Mex Tilapia",
					Headline:    "with Cilantro Lime Couscous and Green Beans",
					Description: "Let’s take tilapia to the next level and turn it into a Tex-Mex-style triumph on your plate. The firm-fleshed fish is given a dusting of panko breadcrumbs and our Southwest spice blend, which ensures that it has satisfying crunch and zesty flavor in every bite. The couscous, green beans, and lime crema on the side come together in a flash, meaning you’ll have this masterpiece of a meal on the table in a matter of minutes.",
					Difficulty:  1,
					PrepTime:    "PT30M",
					ImageLink:   "https://d3hvwccx09j84u.cloudfront.net/0,0/image/5a958c0d30006c33ca2850f2-c352c2d5.jpg",
					Ingredients: TestIngredients,
				},
				&Recipe{
					ID:          23,
					Name:        "Meatloaf à La Mom",
					Headline:    "with Roasted Root Veggies and Thyme Gravy",
					Description: "It’s commonly agreed that no one makes meatloaf quite like Mom. But our chefs will settle for second best by taking a page from her book and passing on this recipe that stays about as true to tradition as can be. These ground beef mini loaves are brushed with a ketchup glaze, served with tender roasted carrots and potatoes, and drizzled with an herby thyme gravy. If you’re a mom (or pop) who wants to claim this recipe as your own, don’t worry—we’ll stay mum.",
					Difficulty:  1,
					PrepTime:    "PT20M",
					ImageLink:   "https://d3hvwccx09j84u.cloudfront.net/0,0/image/5a8f0db7ae08b52cf0617622-a8d742cd.jpg",
					Ingredients: TestIngredients,
				},
			}
			RecipesCon = append(RecipesCon, recipes...)
			sort.Sort(RecipesCon)
			Expect(RecipesCon[0].ID).To(Equal(23))
			Expect(RecipesCon[0].PrepTime).To(Equal("PT20M"))
			Expect(RecipesCon[1].ID).To(Equal(3))
			Expect(RecipesCon[1].PrepTime).To(Equal("PT30M"))
			Expect(RecipesCon[2].ID).To(Equal(1))
			Expect(RecipesCon[2].PrepTime).To(Equal("PT35M"))
			Expect(RecipesCon[3].ID).To(Equal(14))
			Expect(RecipesCon[3].PrepTime).To(Equal("PT45M"))
		})
	})
})
