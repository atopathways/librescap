// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://scap.nist.gov/schema/ocil/2.0
package inter

import (
	"encoding/xml"
)

// Element
type Ocil struct {
	XMLName xml.Name `xml:ocil`

	Generator GeneratorType `xml:"generator"`

	Document *DocumentType `xml:"document"`

	Questionnaires QuestionnairesType `xml:"questionnaires"`

	TestActions TestActionsType `xml:"test_actions"`

	Questions QuestionsType `xml:"questions"`

	Artifacts *ArtifactsType `xml:"artifacts"`

	Variables *VariablesType `xml:"variables"`

	Results *ResultsType `xml:"results"`
}

// Element
type TestAction struct {
	XMLName xml.Name `xml:test_action`

	Revision string `xml:"revision,attr"`

	Notes []string `xml:"notes"`
}

// Element
type QuestionTestAction struct {
	XMLName xml.Name `xml:question_test_action`

	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

// Element
type BooleanQuestionTestAction struct {
	XMLName xml.Name `xml:boolean_question_test_action`

	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	WhenTrue TestActionConditionType `xml:"when_true"`

	WhenFalse TestActionConditionType `xml:"when_false"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

// Element
type ChoiceQuestionTestAction struct {
	XMLName xml.Name `xml:choice_question_test_action`

	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	WhenChoice []ChoiceTestActionConditionType `xml:"when_choice"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

// Element
type NumericQuestionTestAction struct {
	XMLName xml.Name `xml:numeric_question_test_action`

	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

// Element
type StringQuestionTestAction struct {
	XMLName xml.Name `xml:string_question_test_action`

	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	WhenPattern []PatternTestActionConditionType `xml:"when_pattern"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

// Element
type Question struct {
	XMLName xml.Name `xml:question`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

// Element
type BooleanQuestion struct {
	XMLName xml.Name `xml:boolean_question`

	DefaultAnswer string `xml:"default_answer,attr"`

	Model string `xml:"model,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

// Element
type ChoiceQuestion struct {
	XMLName xml.Name `xml:choice_question`

	DefaultAnswerRef string `xml:"default_answer_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

// Element
type NumericQuestion struct {
	XMLName xml.Name `xml:numeric_question`

	DefaultAnswer string `xml:"default_answer,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

// Element
type StringQuestion struct {
	XMLName xml.Name `xml:string_question`

	DefaultAnswer string `xml:"default_answer,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

// Element
type Variable struct {
	XMLName xml.Name `xml:variable`

	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`
}

// Element
type ConstantVariable struct {
	XMLName xml.Name `xml:constant_variable`

	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Value string `xml:"value"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`
}

// Element
type LocalVariable struct {
	XMLName xml.Name `xml:local_variable`

	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Set string `xml:"set"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`
}

// Element
type ExternalVariable struct {
	XMLName xml.Name `xml:external_variable`

	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`
}

// Element
type Target struct {
	XMLName xml.Name `xml:target`

	Revision string `xml:"revision,attr"`

	Name string `xml:"name"`

	Notes []string `xml:"notes"`
}

// Element
type User struct {
	XMLName xml.Name `xml:user`

	Revision string `xml:"revision,attr"`

	Organization []string `xml:"organization"`

	Position []string `xml:"position"`

	Email []string `xml:"email"`

	Name string `xml:"name"`

	Notes []string `xml:"notes"`
}

// Element
type System struct {
	XMLName xml.Name `xml:system`

	Revision string `xml:"revision,attr"`

	Organization string `xml:"organization"`

	Ipaddress []string `xml:"ipaddress"`

	Description *TextType `xml:"description"`

	Name string `xml:"name"`

	Notes []string `xml:"notes"`
}

// Element
type QuestionResult struct {
	XMLName xml.Name `xml:question_result`

	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr"`
}

// Element
type BooleanQuestionResult struct {
	XMLName xml.Name `xml:boolean_question_result`

	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr"`

	Answer bool `xml:"answer"`
}

// Element
type ChoiceQuestionResult struct {
	XMLName xml.Name `xml:choice_question_result`

	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr"`

	Answer ChoiceAnswerType `xml:"answer"`
}

// Element
type NumericQuestionResult struct {
	XMLName xml.Name `xml:numeric_question_result`

	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr"`

	Answer float64 `xml:"answer"`
}

// Element
type StringQuestionResult struct {
	XMLName xml.Name `xml:string_question_result`

	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr"`

	Answer string `xml:"answer"`
}

// Element
type ArtifactValue struct {
	XMLName xml.Name `xml:artifact_value`
}

// Element
type TextArtifactValue struct {
	XMLName xml.Name `xml:text_artifact_value`

	MimeType string `xml:"mime_type,attr"`

	Data string `xml:"data"`
}

// Element
type BinaryArtifactValue struct {
	XMLName xml.Name `xml:binary_artifact_value`

	MimeType string `xml:"mime_type,attr"`

	Data string `xml:"data"`
}

// Element
type ReferenceArtifactValue struct {
	XMLName xml.Name `xml:reference_artifact_value`

	Reference Reference `xml:"reference"`
}

// Element
type Expression struct {
	XMLName xml.Name `xml:expression`

	Value string `xml:"value"`
}

// Element
type WhenPattern struct {
	XMLName xml.Name `xml:when_pattern`

	Pattern string `xml:"pattern,attr"`

	Value string `xml:"value"`
}

// Element
type WhenChoice struct {
	XMLName xml.Name `xml:when_choice`

	ChoiceRef string `xml:"choice_ref,attr"`

	Value string `xml:"value"`
}

// Element
type WhenRange struct {
	XMLName xml.Name `xml:when_range`

	Min string `xml:"min,attr"`

	Max string `xml:"max,attr"`

	Value string `xml:"value"`
}

// Element
type WhenBoolean struct {
	XMLName xml.Name `xml:when_boolean`

	Value string `xml:"value,attr"`

	ValueElm string `xml:"value"`
}

// Element
type Reference struct {
	XMLName xml.Name `xml:reference`

	Href string `xml:"href,attr"`
}

// XSD ComplexType declarations

type OCILType struct {
	Generator GeneratorType `xml:"generator"`

	Document *DocumentType `xml:"document"`

	Questionnaires QuestionnairesType `xml:"questionnaires"`

	TestActions TestActionsType `xml:"test_actions"`

	Questions QuestionsType `xml:"questions"`

	Artifacts *ArtifactsType `xml:"artifacts"`

	Variables *VariablesType `xml:"variables"`

	Results *ResultsType `xml:"results"`
}

type QuestionnairesType struct {
	Questionnaire []QuestionnaireType `xml:"questionnaire"`
}

type QuestionnaireType struct {
	Id string `xml:"id,attr"`

	ChildOnly string `xml:"child_only,attr"`

	Revision string `xml:"revision,attr"`

	Title *TextType `xml:"title"`

	Description *TextType `xml:"description"`

	References *ReferencesType `xml:"references"`

	Actions OperationType `xml:"actions"`

	Notes []string `xml:"notes"`
}

type GeneratorType struct {
	ProductName string `xml:"product_name"`

	ProductVersion string `xml:"product_version"`

	Author []UserType `xml:"author"`

	SchemaVersion float64 `xml:"schema_version"`

	Timestamp string `xml:"timestamp"`

	AdditionalData *ExtensionContainerType `xml:"additional_data"`
}

type ExtensionContainerType struct {
}

type DocumentType struct {
	Title string `xml:"title"`

	Description []string `xml:"description"`

	Notice []string `xml:"notice"`
}

type TestActionsType struct {
	TestAction []TestAction `xml:"test_action"`
}

type QuestionTestActionType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

type BooleanQuestionTestActionType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	WhenTrue TestActionConditionType `xml:"when_true"`

	WhenFalse TestActionConditionType `xml:"when_false"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

type ChoiceQuestionTestActionType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	WhenChoice []ChoiceTestActionConditionType `xml:"when_choice"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

type NumericQuestionTestActionType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

type StringQuestionTestActionType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	WhenPattern []PatternTestActionConditionType `xml:"when_pattern"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

type TestActionRefType struct {
	Negate string `xml:"negate,attr"`

	Text string `xml:",chardata"`
}

type ChoiceTestActionConditionType struct {
	ChoiceRef []string `xml:"choice_ref"`

	ArtifactRefs *ArtifactRefsType `xml:"artifact_refs"`

	Result string `xml:"result"`

	TestActionRef TestActionRefType `xml:"test_action_ref"`
}

type EqualsTestActionConditionType struct {
	VarRef string `xml:"var_ref,attr"`

	Value []float64 `xml:"value"`

	ArtifactRefs *ArtifactRefsType `xml:"artifact_refs"`

	Result string `xml:"result"`

	TestActionRef TestActionRefType `xml:"test_action_ref"`
}

type RangeTestActionConditionType struct {
	Range []RangeType `xml:"range"`

	ArtifactRefs *ArtifactRefsType `xml:"artifact_refs"`

	Result string `xml:"result"`

	TestActionRef TestActionRefType `xml:"test_action_ref"`
}

type PatternTestActionConditionType struct {
	Pattern []PatternType `xml:"pattern"`

	ArtifactRefs *ArtifactRefsType `xml:"artifact_refs"`

	Result string `xml:"result"`

	TestActionRef TestActionRefType `xml:"test_action_ref"`
}

type PatternType struct {
	VarRef string `xml:"var_ref,attr"`

	Text string `xml:",chardata"`
}

type RangeType struct {
	Min *RangeValueType `xml:"min"`

	Max *RangeValueType `xml:"max"`
}

type TestActionConditionType struct {
	ArtifactRefs *ArtifactRefsType `xml:"artifact_refs"`

	Result string `xml:"result"`

	TestActionRef TestActionRefType `xml:"test_action_ref"`
}

type RangeValueType struct {
	Inclusive string `xml:"inclusive,attr"`

	VarRef string `xml:"var_ref,attr"`

	Text string `xml:",chardata"`
}

type QuestionsType struct {
	Question []Question `xml:"question"`

	ChoiceGroup []ChoiceGroupType `xml:"choice_group"`
}

type QuestionTextType struct {
	Sub []SubstitutionTextType `xml:"sub"`

	InnerXml string `xml:",innerxml"`
}

type QuestionType struct {
	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

type BooleanQuestionType struct {
	DefaultAnswer string `xml:"default_answer,attr"`

	Model string `xml:"model,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

type ChoiceQuestionType struct {
	DefaultAnswerRef string `xml:"default_answer_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

type NumericQuestionType struct {
	DefaultAnswer string `xml:"default_answer,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

type StringQuestionType struct {
	DefaultAnswer string `xml:"default_answer,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

type ChoiceType struct {
	Id string `xml:"id,attr"`

	VarRef string `xml:"var_ref,attr"`

	Text string `xml:",chardata"`
}

type ChoiceGroupType struct {
	Id string `xml:"id,attr"`

	Choice []ChoiceType `xml:"choice"`
}

type InstructionsType struct {
	Title TextType `xml:"title"`

	Step []StepType `xml:"step"`
}

type ResultsType struct {
	StartTime string `xml:"start_time,attr"`

	EndTime string `xml:"end_time,attr"`

	Title *TextType `xml:"title"`

	QuestionnaireResults *QuestionnaireResultsType `xml:"questionnaire_results"`

	TestActionResults *TestActionResultsType `xml:"test_action_results"`

	QuestionResults *QuestionResultsType `xml:"question_results"`

	ArtifactResults *ArtifactResultsType `xml:"artifact_results"`

	Targets *TargetsType `xml:"targets"`
}

type QuestionnaireResultsType struct {
	QuestionnaireResult []QuestionnaireResultType `xml:"questionnaire_result"`
}

type TestActionResultsType struct {
	TestActionResult []TestActionResultType `xml:"test_action_result"`
}

type QuestionResultsType struct {
	QuestionResult []QuestionResult `xml:"question_result"`
}

type QuestionnaireResultType struct {
	QuestionnaireRef string `xml:"questionnaire_ref,attr"`

	Result string `xml:"result,attr"`

	ArtifactResults *ArtifactResultsType `xml:"artifact_results"`
}

type TestActionResultType struct {
	TestActionRef string `xml:"test_action_ref,attr"`

	Result string `xml:"result,attr"`

	ArtifactResults *ArtifactResultsType `xml:"artifact_results"`
}

type QuestionResultType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr"`
}

type BooleanQuestionResultType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr"`

	Answer bool `xml:"answer"`
}

type ChoiceQuestionResultType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr"`

	Answer ChoiceAnswerType `xml:"answer"`
}

type NumericQuestionResultType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr"`

	Answer float64 `xml:"answer"`
}

type StringQuestionResultType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr"`

	Answer string `xml:"answer"`
}

type ChoiceAnswerType struct {
	ChoiceRef string `xml:"choice_ref,attr"`
}

type ArtifactsType struct {
	Artifact []ArtifactType `xml:"artifact"`
}

type ArtifactType struct {
	Id string `xml:"id,attr"`

	Persistent string `xml:"persistent,attr"`

	Revision string `xml:"revision,attr"`

	Title TextType `xml:"title"`

	Description TextType `xml:"description"`

	Notes []string `xml:"notes"`
}

type ArtifactRefsType struct {
	ArtifactRef []ArtifactRefType `xml:"artifact_ref"`
}

type ArtifactRefType struct {
	Idref string `xml:"idref,attr"`

	Required string `xml:"required,attr"`
}

type ArtifactResultsType struct {
	ArtifactResult []ArtifactResultType `xml:"artifact_result"`
}

type ArtifactValueType struct {
}

type EmbeddedArtifactValueType struct {
	MimeType string `xml:"mime_type,attr"`
}

type TextArtifactValueType struct {
	MimeType string `xml:"mime_type,attr"`

	Data string `xml:"data"`
}

type BinaryArtifactValueType struct {
	MimeType string `xml:"mime_type,attr"`

	Data string `xml:"data"`
}

type ReferenceArtifactValueType struct {
	Reference Reference `xml:"reference"`
}

type ArtifactResultType struct {
	ArtifactRef string `xml:"artifact_ref,attr"`

	Timestamp string `xml:"timestamp,attr"`

	ArtifactValue ArtifactValue `xml:"artifact_value"`

	Provider string `xml:"provider"`

	Submitter UserType `xml:"submitter"`
}

type TargetsType struct {
	Target []Target `xml:"target"`
}

type UserType struct {
	Revision string `xml:"revision,attr"`

	Organization []string `xml:"organization"`

	Position []string `xml:"position"`

	Email []string `xml:"email"`

	Name string `xml:"name"`

	Notes []string `xml:"notes"`
}

type SystemTargetType struct {
	Revision string `xml:"revision,attr"`

	Organization string `xml:"organization"`

	Ipaddress []string `xml:"ipaddress"`

	Description *TextType `xml:"description"`

	Name string `xml:"name"`

	Notes []string `xml:"notes"`
}

type VariablesType struct {
	Variable []Variable `xml:"variable"`
}

type VariableType struct {
	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`
}

type ConstantVariableType struct {
	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Value string `xml:"value"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`
}

type LocalVariableType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Set string `xml:"set"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`
}

type ExternalVariableType struct {
	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`
}

type SetExpressionBaseType struct {
	Value string `xml:"value"`
}

type SetExpressionPatternType struct {
	Pattern string `xml:"pattern,attr"`

	Value string `xml:"value"`
}

type SetExpressionChoiceType struct {
	ChoiceRef string `xml:"choice_ref,attr"`

	Value string `xml:"value"`
}

type SetExpressionRangeType struct {
	Min string `xml:"min,attr"`

	Max string `xml:"max,attr"`

	Value string `xml:"value"`
}

type SetExpressionBooleanType struct {
	Value string `xml:"value,attr"`

	ValueElm string `xml:"value"`
}

type VariableSetType struct {
	Expression []Expression `xml:"expression"`
}

type SubstitutionTextType struct {
	VarRef string `xml:"var_ref,attr"`
}

type ReferenceType struct {
	Href string `xml:"href,attr"`

	XmlLang string `xml:"lang,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type StepType struct {
	IsDone string `xml:"is_done,attr"`

	IsRequired string `xml:"is_required,attr"`

	Description *TextType `xml:"description"`

	Reference []ReferenceType `xml:"reference"`

	Step []StepType `xml:"step"`
}

type ItemBaseType struct {
	Revision string `xml:"revision,attr"`

	Notes []string `xml:"notes"`
}

type NamedItemBaseType struct {
	Revision string `xml:"revision,attr"`

	Name string `xml:"name"`

	Notes []string `xml:"notes"`
}

type CompoundTestActionType struct {
	Revision string `xml:"revision,attr"`

	Title *TextType `xml:"title"`

	Description *TextType `xml:"description"`

	References *ReferencesType `xml:"references"`

	Actions OperationType `xml:"actions"`

	Notes []string `xml:"notes"`
}

type ReferencesType struct {
	Reference []ReferenceType `xml:"reference"`
}

type OperationType struct {
	Operation string `xml:"operation,attr"`

	Negate string `xml:"negate,attr"`

	TestActionRef []TestActionRefType `xml:"test_action_ref"`
}

type TextType struct {
	XmlLang string `xml:"lang,attr"`

	Text string `xml:",chardata"`
}
