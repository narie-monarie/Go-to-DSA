public class Solution {
    public string DefangIPaddr(string address) {
        var newaddress = address.Replace(".","[.]");
        return newaddress;
    }
}