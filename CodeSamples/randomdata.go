package main

import "fmt"
import "github.com/Pallinder/go-randomdata"

func main() {

	// Print a random silly name
	fmt.Println("A random silly name:", randomdata.SillyName())

	// Print a male title
	fmt.Println("A random male title:", randomdata.Title(randomdata.Male))

	// Print a female title
	fmt.Println("A random female title:", randomdata.Title(randomdata.Female))

	// Print a title with random gender
	fmt.Println("A a random title with gender:", randomdata.Title(randomdata.RandomGender))

	// Print a male first name
	fmt.Println("A random male firstname:", randomdata.FirstName(randomdata.Male))

	// Print a female first name
	fmt.Println("A random female firstname:", randomdata.FirstName(randomdata.Female))

	// Print a last name
	fmt.Println("A random lastname:", randomdata.LastName())

	// Print a male name
	fmt.Println("A random male fullname:", randomdata.FullName(randomdata.Male))

	// Print a female name
	fmt.Println("A random female fullname:", randomdata.FullName(randomdata.Female))

	// Print a name with random gender
	fmt.Println("A random fullname with gender:", randomdata.FullName(randomdata.RandomGender))

	// Print an email
	fmt.Println("A random email:", randomdata.Email())

	// Print a country with full text representation
	fmt.Println("A random country:", randomdata.Country(randomdata.FullCountry))

	// Print a country using ISO 3166-1 alpha-2
	fmt.Println("A random 2 char country:", randomdata.Country(randomdata.TwoCharCountry))

	// Print a country using ISO 3166-1 alpha-3
	fmt.Println("A random 3 char country:", randomdata.Country(randomdata.ThreeCharCountry))

	// Print a currency using ISO 4217
	fmt.Println("A currency ISO 4217:", randomdata.Currency())

	// Print the name of a random city
	fmt.Println("A random city:", randomdata.City())

	// Print the name of a random american state
	fmt.Println("A random state:", randomdata.State(randomdata.Large))

	// Print the name of a random american state using two chars
	fmt.Println("A random state:", randomdata.State(randomdata.Small))

	// Print an american sounding street name
	fmt.Println("A random street:", randomdata.Street())

	// Print an american sounding address
	fmt.Println("A random address:\n", randomdata.Address())

	// Print a random number >= 10 and <= 20
	fmt.Println("A random number between 10 and 20:", randomdata.Number(10, 20))

	// Print a number >= 0 and <= 20
	fmt.Println("A random number between 0 and 20:", randomdata.Number(20))

	// Print a random float >= 0 and <= 20 with decimal point 3
	fmt.Println("A decimal number between 0 and 20 with 3 decimal points:", randomdata.Decimal(0, 20, 3))

	// Print a random float >= 10 and <= 20
	fmt.Println("A decimal between 10 and 20:", randomdata.Decimal(10, 20))

	// Print a random float >= 0 and <= 20
	fmt.Println("A decimal number between 0 and 20:", randomdata.Decimal(20))

	// Print a bool
	fmt.Println("A boolean:", randomdata.Boolean())

	// Print a paragraph
	fmt.Println("A random paragraph:", randomdata.Paragraph())

	// Print a postal code
	fmt.Println("Postal code 'SE':", randomdata.PostalCode("SE"))

	// Print a set of 2 random numbers as a string
	fmt.Println("A set of 2 random numbers:", randomdata.StringNumber(2, "-"))

	// Print a set of 2 random 3-Digits numbers as a string
	fmt.Println("A set of 2 3-digits numbers:", randomdata.StringNumberExt(2, "-", 3))

	// Print a valid random IPv4 address
	fmt.Println("A random IPv4:", randomdata.IpV4Address())

	// Print a valid random IPv6 address
	//fmt.Println("A random IPv6:", randomdata.Ipv6Address())

	// Print a day
	fmt.Println("A random day:", randomdata.Day())

	// Print a month
	fmt.Println("A random month:", randomdata.Month())

	// Print full date like Thursday 22 Aug 2016
	fmt.Println("Print fulldate:", randomdata.FullDate())

	// Get a complete and randomised profile of data generally used for users
	// There are many fields in the profile to use check the Profile struct definition in fullprofile.go
	//profile := randomdata.GenerateProfile(randomdata.Male | randomdata.Female | randomdata.RandomGender)
	//fmt.Printf("The new profile's username is: %s and password (md5)", profile.Login.Username, profile.Login.Md5)
}
