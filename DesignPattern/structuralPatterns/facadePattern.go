package structuralPatterns

import "strconv"

// 外观模式
// 结构型设计模式
// 为复杂系统、程序库或框架提供一个简单 （但有限） 的接口，
// 降低了程序的整体复杂度，有助于将不需要的依赖移动到同一个位置。
// 将内部复杂的子系统封装起来，使其只暴露出最核心的功能。

// example:

// TaoFacade 门面结构
// 封装了复杂的子系统：用户、商品、优惠、库存、支付等系统
// 提供了一个简单的接口
type TaoFacade struct {
	userService    *UserService
	productService *ProductService
	couponService  *CouponService
	stockService   *StockService
	paymentService *PaymentService
}

func NewTaoFacade() *TaoFacade {
	return &TaoFacade{
		userService:    &UserService{},
		productService: &ProductService{},
		couponService:  &CouponService{},
		stockService:   &StockService{},
		paymentService: &PaymentService{},
	}
}

// CreateOrder 创建订单
// 创建订单的过程中，需要调用多个子系统的接口，使用外观模式，可以简化调用过程，提高代码的可读性。
func (t *TaoFacade) CreateOrder(userName, productName string, count int) string {
	// 优惠
	couponInfo := t.couponService.getCoupon(userName)
	// 减少库存
	stockInfo := t.stockService.decreaseFor(productName, count)
	// 计算订单价格
	sumPrice := t.productService.getProductPrice(productName) * float64(count)
	// 支付
	payInfo := t.paymentService.pay(sumPrice)
	// 返回订单信息
	return couponInfo + stockInfo + payInfo + t.userService.getUserAddress(userName)
}

// UserService 用户系统
type UserService struct{}

func (u *UserService) getUserAddress(name string) string {
	return name + "user address"
}

// ProductService 商品系统
type ProductService struct{}

func (p *ProductService) getProductPrice(name string) float64 {
	return 100
}

// CouponService 优惠系统
type CouponService struct{}

func (c *CouponService) getCoupon(name string) string {
	return name + "coupon"
}

// StockService 库存系统
type StockService struct{}

func (s *StockService) decreaseFor(productName string, count int) string {
	return productName + "stock decrease" + strconv.Itoa(count)
}

// PaymentService 支付系统
type PaymentService struct{}

func (p *PaymentService) pay(amount float64) string {
	return "pay " + strconv.FormatFloat(amount, 'f', 2, 64)
}
