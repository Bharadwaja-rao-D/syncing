interface Drawing {
    draw_objs: DrawObject[]
}

interface DrawObject{
    shape: Shape,
    location: location
    scale: scale
}

interface location{
    x_cord: number,
    y_cord: number
}

interface scale{
    scale_x: number,
    scale_y: number
}

enum Shape{
    Rect,
    Circle
}
