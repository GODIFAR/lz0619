// 物品
type item struct {
	name string //物品名
}

// 人物
type character struct {
	name     string //人物名
	leftHand *item  //左手的物品
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v picks up a %v\n", c.name, i.name)
	c.leftHand = i
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%v has nothing to give  \n", c.name)
		return
	}
	if to.leftHand == nil {
		fmt.Printf("%v's hands are full\n", to.name)
	}
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v gives %v a %v \n", c.name, to.name, to.leftHand.name)
}

// String（）接口打印名字
func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v is carrying nothing", c.name)
	}
	return fmt.Sprintf("%v is carrying a %v", c.name, c.leftHand.name)
}
func main() {
	//人物 亚瑟
	arthur := &character{name: "Arthur"}

	//物品 灌木
	shrubbery := &item{name: "shrubbery"}
	arthur.pickup(shrubbery)
	//人物 骑士
	knight := &character{name: "Knight"}
	arthur.give(knight)

	fmt.Println(arthur)
	fmt.Println(knight)
}