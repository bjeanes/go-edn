// -*-go-*-

package edn

%% machine ednParser;
%% write data;

func ParseBytes(data []byte) (EDN, error) {
	cs, p, pe := 0, 0, len(data)

	%%{
		edn = "[]";

		main := edn;
		write init;
		write exec;
	}%%

	return []EDN{}, nil
}
