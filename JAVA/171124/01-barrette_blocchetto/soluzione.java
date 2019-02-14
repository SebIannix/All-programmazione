public class soluzione {
	public static void main(String[]args){
		int[] array = argsToInt(args);
	}
	
	public static int[] argsToInt(String[]args){
		int[] array = new int[args.length];
		for(int i = 0; i < args.length; i++)
			array[i] = Integer.parseInt(args[i]);
		return array;
	}
	
	public static int MAX(int[] array){
		int max = 0;
		for(int j = 0; j < array.length; j++)
			if(max < array[j])
				max = array[j]
		return max;
	}
	
}
