// -*-go-*-

package edn

%% machine ednParser;
%% write data;

func ParseBytes(data []byte) (EDN, error) {
	cs, p, pe := 0, 0, len(data)

	%%{
		set    = '#{'  '}';
		map    = '{' '}';
		list   = '(' ')';
		vector = '[' ']';
		string = '"' '"';

		ws = space | ',';

		edn = (
			vector | 
			list |
			map |
			set |
			string
		);

		main := ws* (edn ws*)?;

		write init;
		write exec;
	}%%

	return []EDN{}, nil
}
