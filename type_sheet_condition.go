/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package lark

// SheetRuleType sheet rule 类型
type SheetRuleType string

// SheetRuleTypeContainsBlanks ...
const (
	SheetRuleTypeContainsBlanks    SheetRuleType = "containsBlanks"    // 为空
	SheetRuleTypeNotContainsBlanks SheetRuleType = "notContainsBlanks" // 不为空
	SheetRuleTypeDuplicateValues   SheetRuleType = "duplicateValues"   // 重复值
	SheetRuleTypeUniqueValues      SheetRuleType = "uniqueValues"      // 唯一值
	SheetRuleTypeCellIs            SheetRuleType = "cellIs"            // 限定值范围
	SheetRuleTypeContainsText      SheetRuleType = "containsText"      // 包含内容
	SheetRuleTypeTimePeriod        SheetRuleType = "timePeriod"        // 日期
)

// SheetRuleAttrOperator sheet rule 操作符
type SheetRuleAttrOperator string

// SheetRuleAttrOperatorEqual ...
const (
	// cellIs
	SheetRuleAttrOperatorEqual              SheetRuleAttrOperator = "equal"              // 限定值范围：等于
	SheetRuleAttrOperatorNotEqual           SheetRuleAttrOperator = "notEqual"           // 限定值范围：不等于
	SheetRuleAttrOperatorGreaterThan        SheetRuleAttrOperator = "greaterThan"        // 限定值范围：大于
	SheetRuleAttrOperatorGreaterThanOrEqual SheetRuleAttrOperator = "greaterThanOrEqual" // 限定值范围：大于或等于
	SheetRuleAttrOperatorLessThan           SheetRuleAttrOperator = "lessThan"           // 限定值范围：小于
	SheetRuleAttrOperatorLessThanOrEqual    SheetRuleAttrOperator = "lessThanOrEqual"    // 限定值范围：小于或等于
	SheetRuleAttrOperatorBetween            SheetRuleAttrOperator = "between"            // 限定值范围：介于
	SheetRuleAttrOperatorNotBetween         SheetRuleAttrOperator = "notBetween"         // 限定值范围：未介于

	// containsText
	SheetRuleAttrOperatorContainsText SheetRuleAttrOperator = "containsText" // 包含以下内容：文本包含
	SheetRuleAttrOperatorNotContains  SheetRuleAttrOperator = "notContains"  // 包含以下内容：文本不包含
	SheetRuleAttrOperatorIs           SheetRuleAttrOperator = "is"           // 包含以下内容：文本为，或者日期为
	SheetRuleAttrOperatorBeginsWith   SheetRuleAttrOperator = "beginsWith"   // 包含以下内容：开头为
	SheetRuleAttrOperatorEndsWith     SheetRuleAttrOperator = "endsWith"     // 包含以下内容：结尾为

)

// SheetRuleAttrTimePeriod sheet rule 日期范围
type SheetRuleAttrTimePeriod string

// SheetRuleAttrTimePeriodYesterday ...
const (
	SheetRuleAttrTimePeriodYesterday SheetRuleAttrTimePeriod = "yesterday" // 日期为：昨天
	SheetRuleAttrTimePeriodToday     SheetRuleAttrTimePeriod = "today"     // 日期为：今天
	SheetRuleAttrTimePeriodTomorrow  SheetRuleAttrTimePeriod = "tomorrow"  // 日期为：明天
	SheetRuleAttrTimePeriodLast7Days SheetRuleAttrTimePeriod = "last7Days" // 日期为：最近7天
)

// SheetRuleAttr sheet rule
type SheetRuleAttr struct {
	Operator   SheetRuleAttrOperator   `json:"operator"`
	Formula    []string                `json:"formula"`
	Text       string                  `json:"text"`
	TimePeriod SheetRuleAttrTimePeriod `json:"time_period"`
}
