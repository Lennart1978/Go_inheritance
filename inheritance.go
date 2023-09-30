package main

import "fmt"

type position struct{
    x float64
    y float64
}
func (p *position)move(x, y float64){
    p.x += x
    p.y += y
}
func (p *position)teleport(x, y float64){
    p.x = x
    p.y = y
}
type player struct{
    position
    name string
    health int
    score int
}
func (p player)String()string{
    return fmt.Sprintf("\n%s\nHealth:%d\nScore:%d\nPosition:%.2f, %.2f",p.name, p.health, p.score, p.x, p.y)
}
func (e enemy)String()string{
    return fmt.Sprintf("\nEnemy position: %.2f %.2f", e.x, e.y)
}
type enemy struct{
    position
    health int
}
func (e *enemy)specialmove(x, y float64){
    e.x *= x
    e.y *= y
}
func main() {
    fmt.Println("Inheritance in Go: - Testprogram")
    for i := 0; i <32; i++{
        fmt.Printf("_")
    }
    var lennart player
    var böse enemy
    
    lennart.name = "Lennart"
    lennart.health, böse.health = 100, 100
    lennart.move(12.5, 44.7)
    böse.move(33.0, 9.0)
    fmt.Println(lennart)
    fmt.Println(böse)
    böse.specialmove(6.66, 6.66)
    lennart.teleport(64.0, 128.0) 
    fmt.Println(lennart)
    fmt.Println(böse)
}
