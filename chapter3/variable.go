/* ***************************************************************************************** */
/*                                                                                           */
/*                                                                                           */
/*   variable.go                                                     __    _            __   */
/*                                                       _________ _/ /_  (_)______  __/ /   */
/*   By: rahisul <rahisul@icloud.com>                   / ___/ __ `/ __ \/ / ___/ / / / /    */
/*                                                     / /  / /_/ / / / / (__  ) /_/ / /     */
/*   Created: 2018/03/26 12:53:13 by rahisul          /_/   \__,_/_/ /_/_/____/___,_/_/      */
/*   Updated: 2018/03/26 13:30:08 by rahisul                                                 */
/*                                                                                           */
/* ***************************************************************************************** */

package main

import "fmt"

func main() {

	var message string = "Hello World!"

	// Below We will implement shorthad varibale delcaration
	var FirstName, LastName string = "Audrey", "Winters"

	// After a variable declaration You cannot declare again "NOT PerMITTED"

	fmt.Println(message)
	fmt.Println(FirstName, " ", LastName)

	//	fmt.Println("vim-go")
}
