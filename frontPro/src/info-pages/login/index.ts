export default class Point{
    constructor(x, y) {
      this.x = x
      this.y = y
      this.r = 1 + Math.random() * 2
      this.sx = Math.random() * 2 - 1
      this.sy = Math.random() * 2 - 1
    }
  
    draw(ctx) {
      ctx.beginPath()
      ctx.arc(this.x, this.y, this.r, 0, 2 * Math.PI)
      ctx.closePath()
      ctx.fillStyle = '#fff'
      ctx.fill()
    }
  
    move(w, h) {
      this.x += this.sx
      this.y += this.sy
      if (this.x > w || this.x < 0) this.sx = -this.sx
      if (this.y > h || this.y < 0) this.sy = -this.sy
    }
  
    drawLine(ctx, p) {
      const dx = this.x - p.x
      const dy = this.y - p.y
      const d = Math.sqrt(dx * dx + dy * dy)
      if (d < 100) {
        var alpha = (100 - d) / 300 * 1
        ctx.beginPath()
        ctx.moveTo(this.x, this.y)
        ctx.lineTo(p.x, p.y)
        ctx.closePath()
        ctx.strokeStyle = 'rgba(255, 255, 255, ' + alpha + ')'
        ctx.strokeWidth = 1
        ctx.stroke()
      }
    }
  }
  