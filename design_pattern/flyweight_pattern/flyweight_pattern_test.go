/**
 * @Author: Sean
 * @Date: 2022/5/31 15:19
 */

package flyweight_pattern

import (
	"fmt"
	"testing"
)

func TestShapeFactory_GetCircle(t *testing.T) {
	shapeF := new(ShapeFactory)
	shape := shapeF.GetCircle("red")

	if _, ok := shapeF.circleMap["red"]; !ok {
		t.Error("map empty,should be 1")
	}

	circle := shape.(*Circle)
	fmt.Println(circle.color)
	if circle.color != "red" {
		t.Error("expected red")
	}
}
