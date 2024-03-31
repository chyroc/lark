package card

import "encoding/json"

func Chart(text string) *ComponentChart {
	return &ComponentChart{}
}

// ComponentChart 图表
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/chart
//
// 飞书卡片提供的图表组件基于 VChart 的图表定义, 支持折线图、面积图、柱状图、饼图、词云等多种数据呈现方式, 帮助你可视化各类信息, 提高信息沟通效率。
// https://www.visactor.io/
//
// 注意事项
// 图表组件仅支持通过撰写卡片 JSON 代码的方式使用, 暂不支持在卡片搭建工具上构建使用。
// 图表组件支持飞书 V7.1 及以上版本的客户端。在低于该版本的飞书客户端上, 图表的内容将展示为一句“请升级客户端为最新版本后查看图表”的占位图。
// 在飞书 7.1 - 7.9 版本上, 图表组件支持的 VChart 版本为 1.6.x; 在飞书 7.10 及以上版本, 图表组件支持的 VChart 版本为 1.8.3。了解 VChart 版本更新, 参考 VChart Changelogs。
// 图表组件暂不支持 JavaScript 语法。
//
// 嵌套规则
// 图表组件可在卡片根节点下或嵌套在分栏容器和折叠面板中使用。
//
// 功能特性
// 基于图表组件绘制的图表, 支持以下功能:
// 图表可交互: 用户可通过点击图表展示数据标签、点击图例实现数据过滤、拖拽缩略轴进行数据筛选。
// 样式自适应: 支持图表多种样式的呈现, 并在不同设备端、不同色彩模式下有良好的自适应展示效果;
// 支持放大查看: PC 端上, 图表支持独立窗口查看; 移动端上, 图表支持点击后全屏查看。
//
// 图表组件基于 VChart 1.6.x 版本, 当前支持折线图、面积图、柱状图、条形图等 13 种图表。本小节列出各个图表的卡片效果和 JSON 示例。要查看各类图表属性的详细说明, 参考 VChart 配置文档。
//
//go:generate generate_set_attrs -type=ComponentChart
//go:generate generate_to_map -type=ComponentChart
type ComponentChart struct {
	componentBase

	// 图表的宽高比。支持以下比例:
	//
	// 1:1
	// 2:1
	// 4:3
	// 16:9
	AspectRatio AspectRatio `json:"aspect_ratio,omitempty"`

	// 图表的主题样式。当图表内存在多个颜色时, 可使用该字段调整颜色样式。若你在 chart_spec 字段中声明了样式类属性, 该字段无效。
	//
	// brand: 默认样式, 与飞书客户端主题样式一致。
	// rainbow: 同色系彩虹色。
	// complementary: 互补色。
	// converse: 反差色。
	// primary: 主色。
	ColorTheme ColorTheme `json:"color_theme,omitempty"`

	// 基于 VChart 的图表定义。详细用法参考 VChart 官方文档。
	// https://www.visactor.io/vchart/guide/tutorial_docs/Chart_Concepts/Understanding_VChart
	//
	// 提示:
	//
	// 飞书 7.1 - 7.9 版本支持的 VChart 版本为 1.6.x;
	// 飞书 7.10 及以上版本支持的 VChart 版本为 1.8.3。
	// todo
	ChartSpec string `json:"chart_spec,omitempty"`

	// 图标是否可在独立窗口查看。可取值:
	//
	// true: 默认值。
	// PC 端: 图表可在独立飞书窗口查看
	// 移动端: 图表可在点击后全屏查看
	// false:
	// PC 端: 图表不支持在独立飞书窗口查看
	// 移动端: 图表不支持在点击后全屏查看
	Preview bool `json:"preview,omitempty"`

	// 图表组件的高度, 可取值:
	//
	// auto: 默认值, 高度将根据宽高比自动计算。
	// [1,999]px: 自定义固定图表高度, 此时宽高比属性 aspect_ratio 失效。
	// 注意: 该属性仅在飞书 7.10 及以上版本生效。
	Height Height `json:"height,omitempty"`
}

type AspectRatio = string

const (
	AspectRatioOneOne      AspectRatio = "1:1"
	AspectRatioTwoOne      AspectRatio = "2:1"
	AspectRatioFourThree   AspectRatio = "4:3"
	AspectRatioSixteenNine AspectRatio = "16:9"
)

type ColorTheme = string

const (
	ColorThemeBrand         ColorTheme = "brand"         // 默认样式, 与飞书客户端主题样式一致。
	ColorThemeRainbow       ColorTheme = "rainbow"       // 同色系彩虹色。
	ColorThemeComplementary ColorTheme = "complementary" // 互补色。
	ColorThemeConverse      ColorTheme = "converse"      // 反差色。
	ColorThemePrimary       ColorTheme = "primary"       // 主色。
)

// MarshalJSON ...
func (r ComponentChart) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "chart"
	return json.Marshal(m)
}

// SetAspectRatio set ComponentChart.AspectRatio attribute
func (r *ComponentChart) SetAspectRatio(val AspectRatio) *ComponentChart {
	r.AspectRatio = val
	return r
}

// SetAspectRatioOneOne set ComponentChart.AspectRatio attribute to AspectRatioOneOne
func (r *ComponentChart) SetAspectRatioOneOne() *ComponentChart {
	r.AspectRatio = AspectRatioOneOne
	return r
}

// SetAspectRatioTwoOne set ComponentChart.AspectRatio attribute to AspectRatioTwoOne
func (r *ComponentChart) SetAspectRatioTwoOne() *ComponentChart {
	r.AspectRatio = AspectRatioTwoOne
	return r
}

// SetAspectRatioFourThree set ComponentChart.AspectRatio attribute to AspectRatioFourThree
func (r *ComponentChart) SetAspectRatioFourThree() *ComponentChart {
	r.AspectRatio = AspectRatioFourThree
	return r
}

// SetAspectRatioSixteenNine set ComponentChart.AspectRatio attribute to AspectRatioSixteenNine
func (r *ComponentChart) SetAspectRatioSixteenNine() *ComponentChart {
	r.AspectRatio = AspectRatioSixteenNine
	return r
}

// SetColorTheme set ComponentChart.ColorTheme attribute
func (r *ComponentChart) SetColorTheme(val ColorTheme) *ComponentChart {
	r.ColorTheme = val
	return r
}

// SetColorThemeBrand set ComponentChart.ColorTheme attribute to ColorThemeBrand
func (r *ComponentChart) SetColorThemeBrand() *ComponentChart {
	r.ColorTheme = ColorThemeBrand
	return r
}

// SetColorThemeRainbow set ComponentChart.ColorTheme attribute to ColorThemeRainbow
func (r *ComponentChart) SetColorThemeRainbow() *ComponentChart {
	r.ColorTheme = ColorThemeRainbow
	return r
}

// SetColorThemeComplementary set ComponentChart.ColorTheme attribute to ColorThemeComplementary
func (r *ComponentChart) SetColorThemeComplementary() *ComponentChart {
	r.ColorTheme = ColorThemeComplementary
	return r
}

// SetColorThemeConverse set ComponentChart.ColorTheme attribute to ColorThemeConverse
func (r *ComponentChart) SetColorThemeConverse() *ComponentChart {
	r.ColorTheme = ColorThemeConverse
	return r
}

// SetColorThemePrimary set ComponentChart.ColorTheme attribute to ColorThemePrimary
func (r *ComponentChart) SetColorThemePrimary() *ComponentChart {
	r.ColorTheme = ColorThemePrimary
	return r
}

// SetChartSpec set ComponentChart.ChartSpec attribute
func (r *ComponentChart) SetChartSpec(val string) *ComponentChart {
	r.ChartSpec = val
	return r
}

// SetPreview set ComponentChart.Preview attribute
func (r *ComponentChart) SetPreview(val bool) *ComponentChart {
	r.Preview = val
	return r
}

// SetHeight set ComponentChart.Height attribute
func (r *ComponentChart) SetHeight(val Height) *ComponentChart {
	r.Height = val
	return r
}

// SetHeightAuto set ComponentChart.Height attribute to HeightAuto
func (r *ComponentChart) SetHeightAuto() *ComponentChart {
	r.Height = HeightAuto
	return r
}

// toMap conv ComponentChart to map
func (r *ComponentChart) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 5)
	if r.AspectRatio != "" {
		res["aspect_ratio"] = r.AspectRatio
	}
	if r.ColorTheme != "" {
		res["color_theme"] = r.ColorTheme
	}
	if r.ChartSpec != "" {
		res["chart_spec"] = r.ChartSpec
	}
	if r.Preview != false {
		res["preview"] = r.Preview
	}
	if r.Height != "" {
		res["height"] = r.Height
	}
	return res
}
