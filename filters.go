package ffmpeg_go

import (
	"fmt"
	"strconv"
)

func AssetType(hasType, expectType string, action string) {
	if hasType != expectType {
		panic(fmt.Sprintf("cannot %s on non-%s", action, expectType))
	}
}

func FilterMultiOutput(streamSpec []*Stream, filterName string, args []string, kwArgs ...KwArgs) *Node {
	return NewFilterNode(filterName, streamSpec, -1, args, MergeKwArgs(kwArgs))
}

func Filter(streamSpec []*Stream, filterName string, args []string, kwArgs ...KwArgs) *Stream {
	return FilterMultiOutput(streamSpec, filterName, args, MergeKwArgs(kwArgs)).Stream("", "")
}

func (s *Stream) Filter(filterName string, args []string, kwArgs ...KwArgs) *Stream {
	AssetType(s.Type, "FilterableStream", "filter")
	return Filter([]*Stream{s}, filterName, args, MergeKwArgs(kwArgs))
}

func (s *Stream) Split() *Node {
	AssetType(s.Type, "FilterableStream", "split")
	return NewFilterNode("split", []*Stream{s}, 1, nil, nil)
}

func (s *Stream) ASplit() *Node {
	AssetType(s.Type, "FilterableStream", "asplit")
	return NewFilterNode("asplit", []*Stream{s}, 1, nil, nil)
}

func (s *Stream) SetPts(expr string) *Node {
	AssetType(s.Type, "FilterableStream", "setpts")
	return NewFilterNode("setpts", []*Stream{s}, 1, []string{expr}, nil)
}

func (s *Stream) Trim(kwargs ...KwArgs) *Stream {
	AssetType(s.Type, "FilterableStream", "trim")
	return NewFilterNode("trim", []*Stream{s}, 1, nil, MergeKwArgs(kwargs)).Stream("", "")
}

func (s *Stream) Overlay(overlayParentNode *Stream, eofAction string, kwargs ...KwArgs) *Stream {
	AssetType(s.Type, "FilterableStream", "overlay")
	if eofAction == "" {
		eofAction = "repeat"
	}
	args := MergeKwArgs(kwargs)
	args["eof_action"] = eofAction
	return NewFilterNode("overlay", []*Stream{s, overlayParentNode}, 2, nil, args).Stream("", "")
}

func (s *Stream) HFlip(kwargs ...KwArgs) *Stream {
	AssetType(s.Type, "FilterableStream", "hflip")
	return NewFilterNode("hflip", []*Stream{s}, 1, nil, MergeKwArgs(kwargs)).Stream("", "")
}

func (s *Stream) VFlip(kwargs ...KwArgs) *Stream {
	AssetType(s.Type, "FilterableStream", "vflip")
	return NewFilterNode("vflip", []*Stream{s}, 1, nil, MergeKwArgs(kwargs)).Stream("", "")
}

func (s *Stream) Crop(x, y, w, h int, kwargs ...KwArgs) *Stream {
	AssetType(s.Type, "FilterableStream", "crop")
	return NewFilterNode("crop", []*Stream{s}, 1, []string{
		strconv.Itoa(w), strconv.Itoa(h), strconv.Itoa(x), strconv.Itoa(y),
	}, MergeKwArgs(kwargs)).Stream("", "")
}

func (s *Stream) DrawBox(x, y, w, h int, color string, thickness int, kwargs ...KwArgs) *Stream {
	AssetType(s.Type, "FilterableStream", "drawbox")
	args := MergeKwArgs(kwargs)
	if thickness != 0 {
		args["t"] = thickness
	}
	return NewFilterNode("drawbox", []*Stream{s}, 1, []string{
		strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(w), strconv.Itoa(h), color,
	}, args).Stream("", "")
}

func (s *Stream) Drawtext(text string, x, y int, escape bool, kwargs ...KwArgs) *Stream {
	AssetType(s.Type, "FilterableStream", "drawtext")
	args := MergeKwArgs(kwargs)
	if escape {
		text = fmt.Sprintf("%q", text)
	}
	if text != "" {
		args["text"] = text
	}
	if x != 0 {
		args["x"] = x
	}

	if y != 0 {
		args["y"] = y
	}

	return NewFilterNode("drawtext", []*Stream{s}, 1, nil, args).Stream("", "")
}

func Concat(streams []*Stream, kwargs ...KwArgs) *Stream {
	args := MergeKwArgs(kwargs)
	vsc := args.GetDefault("v", 1).(int)
	asc := args.GetDefault("a", 0).(int)
	sc := vsc + asc
	if len(streams)%sc != 0 {
		panic("streams count not valid")
	}
	args["n"] = len(streams) / sc
	return NewFilterNode("concat", streams, -1, nil, args).Stream("", "")
}

func (s *Stream) Concat(streams []*Stream, kwargs ...KwArgs) *Stream {
	return Concat(append(streams, s), MergeKwArgs(kwargs))
}

func (s *Stream) ZoomPan(kwargs ...KwArgs) *Stream {
	AssetType(s.Type, "FilterableStream", "zoompan")
	return NewFilterNode("zoompan", []*Stream{s}, 1, nil, MergeKwArgs(kwargs)).Stream("", "")
}

func (s *Stream) Hue(kwargs ...KwArgs) *Stream {
	AssetType(s.Type, "FilterableStream", "hue")
	return NewFilterNode("hue", []*Stream{s}, 1, nil, MergeKwArgs(kwargs)).Stream("", "")
}

func (s *Stream) ColorChannelMixer(kwargs ...KwArgs) *Stream {
	AssetType(s.Type, "FilterableStream", "colorchannelmixer")
	return NewFilterNode("colorchannelmixer", []*Stream{s}, 1, nil, MergeKwArgs(kwargs)).Stream("", "")
}
