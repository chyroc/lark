module.exports = {
  title: 'Lark Go SDK',
  description: 'chyroc/lark 文档',
  base: '/lark/',
  locales: {
    // 键名是该语言所属的子路径
    // 作为特例，默认语言可以使用 '/' 作为其路径。
    '/': {
      lang: 'en-US', // 将会被设置为 <html> 的 lang 属性
    },
    '/zh/': {
      lang: 'zh-CN',
    }
  },
  themeConfig: {
    locales: {
      '/': {
        selectText: 'Languages',
        label: 'English',
        ariaLabel: 'Languages',
        editLinkText: 'Edit this page on GitHub',
        algolia: {},
        nav: [
          { text: 'APIs', link: '/apis/' },
          { text: 'Examples', link: '/examples/' },
          { text: 'GitHub', link: 'https://github.com/chyroc/lark' }
        ]
      },
      '/zh/': {
        // 多语言下拉菜单的标题
        selectText: '选择语言',
        // 该语言在下拉菜单中的标签
        label: '简体中文',
        // 编辑链接文字
        editLinkText: '在 GitHub 上编辑此页',
        // 当前 locale 的 algolia docsearch 选项
        algolia: {},
        nav: [
          { text: '接口', link: '/zh/apis/' },
          { text: '使用示例', link: '/zh/examples/' },
          { text: 'GitHub', link: 'https://github.com/chyroc/lark' }
        ]
      }
    }
  }
}