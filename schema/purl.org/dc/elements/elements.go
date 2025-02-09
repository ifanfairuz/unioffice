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

package elements ;import (_c "encoding/xml";_g "fmt";_e "github.com/ifanfairuz/unioffice";_af "github.com/ifanfairuz/unioffice/common/logger";);func NewElementContainer ()*ElementContainer {_gac :=&ElementContainer {};return _gac };

// ValidateWithPath validates the ElementContainer and its children, prefixing error messages with path
func (_ag *ElementContainer )ValidateWithPath (path string )error {for _faa ,_ge :=range _ag .Choice {if _gg :=_ge .ValidateWithPath (_g .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_faa ));_gg !=nil {return _gg ;
};};return nil ;};type Any struct{SimpleLiteral };

// ValidateWithPath validates the Any and its children, prefixing error messages with path
func (_afg *Any )ValidateWithPath (path string )error {if _cf :=_afg .SimpleLiteral .ValidateWithPath (path );_cf !=nil {return _cf ;};return nil ;};func NewElementsGroup ()*ElementsGroup {_ad :=&ElementsGroup {};return _ad };func NewElementsGroupChoice ()*ElementsGroupChoice {_gd :=&ElementsGroupChoice {};
return _gd };func (_ga *Any )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {return _ga .SimpleLiteral .MarshalXML (e ,start );};

// ValidateWithPath validates the ElementsGroup and its children, prefixing error messages with path
func (_ddf *ElementsGroup )ValidateWithPath (path string )error {for _aff ,_ece :=range _ddf .Choice {if _bfa :=_ece .ValidateWithPath (_g .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_aff ));_bfa !=nil {return _bfa ;
};};return nil ;};func NewAny ()*Any {_cc :=&Any {};_cc .SimpleLiteral =*NewSimpleLiteral ();return _cc };func (_cdd *ElementsGroupChoice )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _cdd .Any !=nil {_deb :=_c .StartElement {Name :_c .Name {Local :"\u0064\u0063\u003a\u0061\u006e\u0079"}};
for _ ,_dg :=range _cdd .Any {e .EncodeElement (_dg ,_deb );};};return nil ;};type SimpleLiteral struct{};func (_fa *ElementContainer )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072";
e .EncodeToken (start );if _fa .Choice !=nil {for _ ,_ca :=range _fa .Choice {_ca .MarshalXML (e ,_c .StartElement {});};};e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};

// Validate validates the ElementsGroup and its children
func (_abg *ElementsGroup )Validate ()error {return _abg .ValidateWithPath ("\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070");};func (_ec *ElementsGroup )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_fbd :for {_dc ,_ebd :=d .Token ();
if _ebd !=nil {return _ebd ;};switch _agg :=_dc .(type ){case _c .StartElement :switch _agg .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_bd :=NewElementsGroupChoice ();
if _dda :=d .DecodeElement (&_bd .Any ,&_agg );_dda !=nil {return _dda ;};_ec .Choice =append (_ec .Choice ,_bd );default:_af .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006de\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070 \u0025\u0076",_agg .Name );
if _db :=d .Skip ();_db !=nil {return _db ;};};case _c .EndElement :break _fbd ;case _c .CharData :};};return nil ;};

// Validate validates the Any and its children
func (_cd *Any )Validate ()error {return _cd .ValidateWithPath ("\u0041\u006e\u0079")};func (_gga *ElementsGroupChoice )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_gdd :for {_eda ,_gdg :=d .Token ();if _gdg !=nil {return _gdg ;};switch _gb :=_eda .(type ){case _c .StartElement :switch _gb .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_edd :=NewAny ();
if _bfd :=d .DecodeElement (_edd ,&_gb );_bfd !=nil {return _bfd ;};_gga .Any =append (_gga .Any ,_edd );default:_af .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072ou\u0070\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076",_gb .Name );
if _bdc :=d .Skip ();_bdc !=nil {return _bdc ;};};case _c .EndElement :break _gdd ;case _c .CharData :};};return nil ;};

// ValidateWithPath validates the ElementsGroupChoice and its children, prefixing error messages with path
func (_cg *ElementsGroupChoice )ValidateWithPath (path string )error {for _fd ,_cfd :=range _cg .Any {if _eec :=_cfd .ValidateWithPath (_g .Sprintf ("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d",path ,_fd ));_eec !=nil {return _eec ;};
};return nil ;};func (_ed *Any )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_ed .SimpleLiteral =*NewSimpleLiteral ();for {_b ,_f :=d .Token ();if _f !=nil {return _g .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0041\u006e\u0079\u003a\u0020\u0025\u0073",_f );
};if _d ,_aa :=_b .(_c .EndElement );_aa &&_d .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the SimpleLiteral and its children, prefixing error messages with path
func (_fac *SimpleLiteral )ValidateWithPath (path string )error {return nil };

// Validate validates the ElementContainer and its children
func (_eee *ElementContainer )Validate ()error {return _eee .ValidateWithPath ("\u0045\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072");};type ElementsGroup struct{Choice []*ElementsGroupChoice ;};type ElementContainer struct{Choice []*ElementsGroupChoice ;
};

// Validate validates the SimpleLiteral and its children
func (_dea *SimpleLiteral )Validate ()error {return _dea .ValidateWithPath ("\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c");};func NewSimpleLiteral ()*SimpleLiteral {_eea :=&SimpleLiteral {};return _eea };func (_fb *ElementContainer )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_eb :for {_fbe ,_fe :=d .Token ();
if _fe !=nil {return _fe ;};switch _ee :=_fbe .(type ){case _c .StartElement :switch _ee .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_cad :=NewElementsGroupChoice ();
if _be :=d .DecodeElement (&_cad .Any ,&_ee );_be !=nil {return _be ;};_fb .Choice =append (_fb .Choice ,_cad );default:_af .Log .Debug ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0025v",_ee .Name );
if _de :=d .Skip ();_de !=nil {return _de ;};};case _c .EndElement :break _eb ;case _c .CharData :};};return nil ;};type ElementsGroupChoice struct{Any []*Any ;};func (_bf *ElementsGroup )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _bf .Choice !=nil {for _ ,_ab :=range _bf .Choice {_ab .MarshalXML (e ,_c .StartElement {});
};};return nil ;};func (_afgb *SimpleLiteral )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_eddc ,_ggc :=d .Token ();if _ggc !=nil {return _g .Errorf ("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0053\u0069\u006d\u0070l\u0065L\u0069t\u0065\u0072\u0061\u006c\u003a\u0020\u0025s",_ggc );
};if _eeg ,_bb :=_eddc .(_c .EndElement );_bb &&_eeg .Name ==start .Name {break ;};};return nil ;};func (_eca *SimpleLiteral )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });
return nil ;};

// Validate validates the ElementsGroupChoice and its children
func (_dba *ElementsGroupChoice )Validate ()error {return _dba .ValidateWithPath ("\u0045\u006c\u0065\u006den\u0074\u0073\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u006f\u0069\u0063\u0065");};func init (){_e .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c",NewSimpleLiteral );
_e .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072",NewElementContainer );
_e .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0061\u006e\u0079",NewAny );_e .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070",NewElementsGroup );
};