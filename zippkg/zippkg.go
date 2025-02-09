//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package zippkg ;import (_b "archive/zip";_cf "bytes";_de "encoding/xml";_bd "fmt";_gg "github.com/ifanfairuz/unioffice";_dd "github.com/ifanfairuz/unioffice/algo";_ed "github.com/ifanfairuz/unioffice/common/tempstorage";_e "github.com/ifanfairuz/unioffice/schema/soo/pkg/relationships";
_bf "io";_a "path";_c "sort";_bff "strings";_d "time";);

// Decode unmarshals the content of a *zip.File as XML to a given destination.
func Decode (f *_b .File ,dest interface{})error {_cab ,_cg :=f .Open ();if _cg !=nil {return _bd .Errorf ("e\u0072r\u006f\u0072\u0020\u0072\u0065\u0061\u0064\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",f .Name ,_cg );};defer _cab .Close ();_bfad :=_de .NewDecoder (_cab );
if _bdg :=_bfad .Decode (dest );_bdg !=nil {return _bd .Errorf ("e\u0072\u0072\u006f\u0072 d\u0065c\u006f\u0064\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",f .Name ,_bdg );};if _dg ,_fg :=dest .(*_e .Relationships );_fg {for _bfbg ,_be :=range _dg .Relationship {switch _be .TypeAttr {case _gg .OfficeDocumentTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .OfficeDocumentType ;
case _gg .StylesTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .StylesType ;case _gg .ThemeTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .ThemeType ;case _gg .ControlTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .ControlType ;case _gg .SettingsTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .SettingsType ;
case _gg .ImageTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .ImageType ;case _gg .CommentsTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .CommentsType ;case _gg .ThumbnailTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .ThumbnailType ;
case _gg .DrawingTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .DrawingType ;case _gg .ChartTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .ChartType ;case _gg .ExtendedPropertiesTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .ExtendedPropertiesType ;
case _gg .CustomXMLTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .CustomXMLType ;case _gg .WorksheetTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .WorksheetType ;case _gg .SharedStringsTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .SharedStringsType ;
case _gg .TableTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .TableType ;case _gg .HeaderTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .HeaderType ;case _gg .FooterTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .FooterType ;case _gg .NumberingTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .NumberingType ;
case _gg .FontTableTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .FontTableType ;case _gg .WebSettingsTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .WebSettingsType ;case _gg .FootNotesTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .FootNotesType ;
case _gg .EndNotesTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .EndNotesType ;case _gg .SlideTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .SlideType ;case _gg .VMLDrawingTypeStrict :_dg .Relationship [_bfbg ].TypeAttr =_gg .VMLDrawingType ;
};};_c .Slice (_dg .Relationship ,func (_fe ,_ddb int )bool {_bde :=_dg .Relationship [_fe ];_cfab :=_dg .Relationship [_ddb ];return _dd .NaturalLess (_bde .IdAttr ,_cfab .IdAttr );});};return nil ;};

// ExtractToDiskTmp extracts a zip file to a temporary file in a given path,
// returning the name of the file.
func ExtractToDiskTmp (f *_b .File ,path string )(string ,error ){_fag ,_ebb :=_ed .TempFile (path ,"\u007a\u007a");if _ebb !=nil {return "",_ebb ;};defer _fag .Close ();_da ,_ebb :=f .Open ();if _ebb !=nil {return "",_ebb ;};defer _da .Close ();_ ,_ebb =_bf .Copy (_fag ,_da );
if _ebb !=nil {return "",_ebb ;};return _fag .Name (),nil ;};func MarshalXMLByTypeIndex (z *_b .Writer ,dt _gg .DocType ,typ string ,idx int ,v interface{})error {_aca :=_gg .AbsoluteFilename (dt ,typ ,idx );return MarshalXML (z ,_aca ,v );};func MarshalXMLByType (z *_b .Writer ,dt _gg .DocType ,typ string ,v interface{})error {_ec :=_gg .AbsoluteFilename (dt ,typ ,0);
return MarshalXML (z ,_ec ,v );};

// MarshalXML creates a file inside of a zip and marshals an object as xml, prefixing it
// with a standard XML header.
func MarshalXML (z *_b .Writer ,filename string ,v interface{})error {_edd :=&_b .FileHeader {};_edd .Method =_b .Deflate ;_edd .Name =filename ;_edd .SetModTime (_d .Now ());_bab ,_bdgf :=z .CreateHeader (_edd );if _bdgf !=nil {return _bd .Errorf ("\u0063\u0072\u0065\u0061ti\u006e\u0067\u0020\u0025\u0073\u0020\u0069\u006e\u0020\u007a\u0069\u0070\u003a\u0020%\u0073",filename ,_bdgf );
};_ ,_bdgf =_bab .Write ([]byte (XMLHeader ));if _bdgf !=nil {return _bd .Errorf ("\u0063\u0072e\u0061\u0074\u0069\u006e\u0067\u0020\u0078\u006d\u006c\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0025\u0073: \u0025\u0073",filename ,_bdgf );
};if _bdgf =_de .NewEncoder (SelfClosingWriter {_bab }).Encode (v );_bdgf !=nil {return _bd .Errorf ("\u006d\u0061\u0072\u0073\u0068\u0061\u006c\u0069\u006e\u0067\u0020\u0025s\u003a\u0020\u0025\u0073",filename ,_bdgf );};_ ,_bdgf =_bab .Write (_dfb );return _bdgf ;
};

// AddFileFromBytes takes a byte array and adds it at a given path to a zip file.
func AddFileFromBytes (z *_b .Writer ,zipPath string ,data []byte )error {_ffa ,_df :=z .Create (zipPath );if _df !=nil {return _bd .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_df );
};_ ,_df =_bf .Copy (_ffa ,_cf .NewReader (data ));return _df ;};

// AddFileFromDisk reads a file from internal storage and adds it at a given path to a zip file.
// TODO: Rename to AddFileFromStorage in next major version release (v2).
// NOTE: If disk storage cannot be used, memory storage can be used instead by calling memstore.SetAsStorage().
func AddFileFromDisk (z *_b .Writer ,zipPath ,storagePath string )error {_ffg ,_dc :=z .Create (zipPath );if _dc !=nil {return _bd .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_dc );
};_ga ,_dc :=_ed .Open (storagePath );if _dc !=nil {return _bd .Errorf ("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",storagePath ,_dc );};defer _ga .Close ();_ ,_dc =_bf .Copy (_ffg ,_ga );return _dc ;
};

// AddTarget allows documents to register decode targets. Path is a path that
// will be found in the zip file and ifc is an XML element that the file will be
// unmarshaled to.  filePath is the absolute path to the target, ifc is the
// object to decode into, sourceFileType is the type of file that the reference
// was discovered in, and index is the index of the source file type.
func (_ddf *DecodeMap )AddTarget (filePath string ,ifc interface{},sourceFileType string ,idx uint32 )bool {if _ddf ._f ==nil {_ddf ._f =make (map[string ]Target );_ddf ._ac =make (map[*_e .Relationships ]string );_ddf ._bfb =make (map[string ]struct{});
_ddf ._gcf =make (map[string ]int );};if _a .IsAbs (filePath ){filePath =_bff .TrimPrefix (filePath ,"\u002f");};_dec :=_a .Clean (filePath );if _ ,_bdc :=_ddf ._bfb [_dec ];_bdc {return false ;};_ddf ._bfb [_dec ]=struct{}{};_ddf ._f [_dec ]=Target {Path :_dec ,Typ :sourceFileType ,Ifc :ifc ,Index :idx };
return true ;};var _dfb =[]byte {'\r','\n'};

// Decode loops decoding targets registered with AddTarget and calling th
func (_aea *DecodeMap )Decode (files []*_b .File )error {_gf :=1;for _gf > 0{for len (_aea ._aee )> 0{_ba :=_aea ._aee [0];_aea ._aee =_aea ._aee [1:];_fd :=_ba .Ifc .(*_e .Relationships );for _ ,_cfa :=range _fd .Relationship {_ff :=_aea ._ac [_fd ];_cfa .TargetAttr =_bff .TrimPrefix (_cfa .TargetAttr ,"\u002f");
if _bff .IndexByte (_ff ,'/')> -1{_eb :=_ff [:_bff .IndexByte (_ff ,'/')+1];if _bff .HasPrefix (_cfa .TargetAttr ,_eb ){_ff ="";};};if _bff .HasPrefix (_cfa .TargetAttr ,_ff ){_ff ="";};_aea ._ce (_aea ,_ff +_cfa .TargetAttr ,_cfa .TypeAttr ,files ,_cfa ,_ba );
};};for _acb ,_bb :=range files {if _bb ==nil {continue ;};if _cff ,_edf :=_aea ._f [_bb .Name ];_edf {delete (_aea ._f ,_bb .Name );if _cb :=Decode (_bb ,_cff .Ifc );_cb !=nil {return _cb ;};files [_acb ]=nil ;if _cag ,_cad :=_cff .Ifc .(*_e .Relationships );
_cad {_aea ._aee =append (_aea ._aee ,_cff );_cc ,_ :=_a .Split (_a .Clean (_bb .Name +"\u002f\u002e\u002e\u002f"));_aea ._ac [_cag ]=_cc ;_gf ++;};};};_gf --;};return nil ;};var _dcf =[]byte {'/','>'};func (_dgf SelfClosingWriter )Write (b []byte )(int ,error ){_ded :=0;
_cd :=0;for _ece :=0;_ece < len (b )-2;_ece ++{if b [_ece ]=='>'&&b [_ece +1]=='<'&&b [_ece +2]=='/'{_gae :=[]byte {};_dgfc :=_ece ;for _gfe :=_ece ;_gfe >=0;_gfe --{if b [_gfe ]==' '{_dgfc =_gfe ;}else if b [_gfe ]=='<'{_gae =b [_gfe +1:_dgfc ];break ;
};};_ea :=[]byte {};for _dfg :=_ece +3;_dfg < len (b );_dfg ++{if b [_dfg ]=='>'{_ea =b [_ece +3:_dfg ];break ;};};if !_cf .Equal (_gae ,_ea ){continue ;};_bef ,_ffgd :=_dgf .W .Write (b [_ded :_ece ]);if _ffgd !=nil {return _cd +_bef ,_ffgd ;};_cd +=_bef ;
_ ,_ffgd =_dgf .W .Write (_dcf );if _ffgd !=nil {return _cd ,_ffgd ;};_cd +=3;for _gdd :=_ece +2;_gdd < len (b )&&b [_gdd ]!='>';_gdd ++{_cd ++;_ded =_gdd +2;_ece =_ded ;};};};_gca ,_bag :=_dgf .W .Write (b [_ded :]);return _gca +_cd ,_bag ;};type Target struct{Path string ;
Typ string ;Ifc interface{};Index uint32 ;};func (_acd *DecodeMap )IndexFor (path string )int {return _acd ._gcf [path ]};const XMLHeader ="\u003c\u003f\u0078\u006d\u006c\u0020\u0076e\u0072\u0073\u0069o\u006e\u003d\u00221\u002e\u0030\"\u0020\u0065\u006e\u0063\u006f\u0064i\u006eg=\u0022\u0055\u0054\u0046\u002d\u0038\u0022\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u006c\u006f\u006e\u0065\u003d\u0022\u0079\u0065\u0073\u0022\u003f\u003e"+"\u000a";


// DecodeMap is used to walk a tree of relationships, decoding files and passing
// control back to the document.
type DecodeMap struct{_f map[string ]Target ;_ac map[*_e .Relationships ]string ;_aee []Target ;_ce OnNewRelationshipFunc ;_bfb map[string ]struct{};_gcf map[string ]int ;};func (_fa *DecodeMap )RecordIndex (path string ,idx int ){_fa ._gcf [path ]=idx };


// SelfClosingWriter wraps a writer and replaces XML tags of the
// type <foo></foo> with <foo/>
type SelfClosingWriter struct{W _bf .Writer ;};

// OnNewRelationshipFunc is called when a new relationship has been discovered.
//
// target is a resolved path that takes into account the location of the
// relationships file source and should be the path in the zip file.
//
// files are passed so non-XML files that can't be handled by AddTarget can be
// decoded directly (e.g. images)
//
// rel is the actual relationship so its target can be modified if the source
// target doesn't match where unioffice will write the file (e.g. read in
// 'xl/worksheets/MyWorksheet.xml' and we'll write out
// 'xl/worksheets/sheet1.xml')
type OnNewRelationshipFunc func (_ca *DecodeMap ,_bfa ,_gd string ,_gc []*_b .File ,_bc *_e .Relationship ,_ae Target )error ;

// SetOnNewRelationshipFunc sets the function to be called when a new
// relationship has been discovered.
func (_ag *DecodeMap )SetOnNewRelationshipFunc (fn OnNewRelationshipFunc ){_ag ._ce =fn };

// RelationsPathFor returns the relations path for a given filename.
func RelationsPathFor (path string )string {_bbg :=_bff .Split (path ,"\u002f");_cbe :=_bff .Join (_bbg [0:len (_bbg )-1],"\u002f");_bbga :=_bbg [len (_bbg )-1];_cbe +="\u002f_\u0072\u0065\u006c\u0073\u002f";_bbga +="\u002e\u0072\u0065l\u0073";return _cbe +_bbga ;
};