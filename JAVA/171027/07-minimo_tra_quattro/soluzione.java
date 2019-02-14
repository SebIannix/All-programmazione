import java.util.*;
public class soluzione {
	public static void main(String[]args) {
		Scanner in = new Scanner(System.in);
		
		int a = in.nextInt();
		int b = in.nextInt();
		int c = in.nextInt();
		int d = in.nextInt();
		int min = a;
		
		
		if (b < a & b < c & b < d)
			min = b;
		if (c < a & c < b & c < d)
			min = c;
		if (d < a & d < b & d < c)
			min = d;
			
		System.out.println(min);
		
	}
}
