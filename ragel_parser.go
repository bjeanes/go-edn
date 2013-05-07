
// line 1 "parser.rl"
// -*-go-*-

package edn


// line 6 "parser.rl"

// line 11 "ragel_parser.go"
var _ednParser_key_offsets []byte = []byte{
	0, 0, 1, 2, 
}

var _ednParser_trans_keys []byte = []byte{
	91, 93, 
}

var _ednParser_single_lengths []byte = []byte{
	0, 1, 1, 0, 
}

var _ednParser_range_lengths []byte = []byte{
	0, 0, 0, 0, 
}

var _ednParser_index_offsets []byte = []byte{
	0, 0, 2, 4, 
}

var _ednParser_trans_targs []byte = []byte{
	2, 0, 3, 0, 0, 
}

const ednParser_start int = 1
const ednParser_first_final int = 3
const ednParser_error int = 0

const ednParser_en_main int = 1


// line 7 "parser.rl"

func ParseBytes(data []byte) (EDN, error) {
	cs, p, pe := 0, 0, len(data)

    
// line 49 "ragel_parser.go"
	{
	cs = ednParser_start
	}

// line 54 "ragel_parser.go"
	{
	var _klen int
	var _trans int
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_ednParser_key_offsets[cs])
	_trans = int(_ednParser_index_offsets[cs])

	_klen = int(_ednParser_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _ednParser_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _ednParser_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_ednParser_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _ednParser_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _ednParser_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	cs = int(_ednParser_trans_targs[_trans])

	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	_out: {}
	}

// line 17 "parser.rl"


	return []EDN{}, nil
}
