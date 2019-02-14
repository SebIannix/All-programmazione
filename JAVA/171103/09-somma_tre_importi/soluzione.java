import java.util.*;
import java.math.*;
public class soluzione {
	public static void main(String[]args) {
		Scanner in = new Scanner(System.in);
		
		String s = in.next();
		String r = in.next();
		String t = in.next();
		String sum = s + r + t;
		BigDecimal k = new BigDecimal(sum);
		
		System.out.println(sum);
		
	}
}
