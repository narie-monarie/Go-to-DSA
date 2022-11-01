class Solution {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        char[] chars = s.toCharArray();
        for(char a : chars){
            if(a == '(' || a == '[' || a == '{'){
                stack.push(a);
                continue;
            }else if(stack.empty())
                return false;
            char top = stack.pop();
            
            if(top == '(' && a != ')')
                return false;
            else if(top == '[' && a != ']')
                return false;
            if(top == '{' && a != '}')
                return false;
        }
        return stack.empty();
    }
}