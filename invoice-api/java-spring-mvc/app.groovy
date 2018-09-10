@RestController
class InvoiceMe {
	@RequestMapping("/invoice")
	String home() {
		return "Hello World";
	}
}