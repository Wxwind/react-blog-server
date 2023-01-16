# 在 vue3-ts-vite 中使用 svg 图片

## 前言

学习 vue3 的时候觉得 element plus 提供的图标太少，想要使用自己的 svg 图片，去网上搜罗了一大筐教程，要么循环引用要么还有一大堆问题，最后总算弄成功了，于是做个记录给自己和大家做个参考

## 步骤

1.安装 svg-sprite-loader,这里使用的是 6.0.11 版本

> npm install svg-sprite-loader

2.项目的 svg 图片存放在 src/icons/svg 下，我们在这里创建两个文件 svgIcon.ts 和 svgIcon.vue（在哪创建和文件名字并没有任何要求）

![image-20221117145717884](https://c2.im5i.com/2023/01/10/YPSmG.png)

3.在 svgIcon.ts 中加入下列代码(如果报错找不到 fs 模块请跳转到[附录 1-导入@type/node](#import-type-node)

```typescript
import { readFileSync, readdirSync } from "fs";

let idPerfix = "";
const svgTitle = /<svg([^>+].*?)>/;
const clearHeightWidth = /(width|height)="([^>+].*?)"/g;
const hasViewBox = /(viewBox="[^>+].*?")/g;
const clearReturn = /(\r)|(\n)/g;

// 查找svg文件
function svgFind(e: any): any {
  const arr = [];
  const dirents = readdirSync(e, { withFileTypes: true });
  for (const dirent of dirents) {
    if (dirent.isDirectory()) arr.push(...svgFind(e + dirent.name + "/"));
    else {
      const svg = readFileSync(e + dirent.name)
        .toString()
        .replace(clearReturn, "")
        .replace(svgTitle, ($1, $2) => {
          let width = 0,
            height = 0,
            content = $2.replace(clearHeightWidth, (s1: any, s2: any, s3: any) => {
              if (s2 === "width") width = s3;
              else if (s2 === "height") height = s3;
              return "";
            });
          if (!hasViewBox.test($2)) content += `viewBox="0 0 ${width} ${height}"`;
          return `<symbol id="${idPerfix}-${dirent.name.replace(".svg", "")}" ${content}>`;
        })
        .replace("</svg>", "</symbol>");
      arr.push(svg);
    }
  }
  return arr;
}

// 生成svg
export const createSvg = (path: any, perfix = "icon") => {
  if (path === "") return;
  idPerfix = perfix;
  const res = svgFind(path);
  return {
    name: "svg-transform",
    transformIndexHtml(dom: String) {
      return dom.replace(
        "<body>",
        `<body><svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" style="position: absolute; width: 0; height: 0">${res.join(
          ""
        )}</svg>`
      );
    },
  };
};
```

4.在 svgIcon.vue 中加入下列代码

```vue
<template>
  <svg :class="svgClass" v-bind="$attrs" :style="{ color: color, fontSize: size }">
    <use :href="iconName"></use>
  </svg>
</template>

<script setup lang="ts">
import { computed } from "vue";
const props = defineProps({
  name: {
    type: String,
    required: true,
  },
  color: {
    type: String,
    default: "",
  },
  size: {
    type: Number,
    default: "",
  },
});
const iconName = computed(() => `#icon-${props.name}`);
const svgClass = computed(() => {
  if (props.name) return `svg-icon icon-${props.name}`;
  return "svg-icon";
});
</script>

<style scoped>
.svg-icon {
  width: 1em;
  height: 1em;
  fill: currentColor;
  vertical-align: middle;
}
</style>
```

5.在 src/main.ts 导入组件

```typescript
import { createApp } from "vue";
import App from "./App.vue";

import svgIcon from "./icons/svgIcon.vue";

createApp(App).component("svg-icon", svgIcon).mount("#app");
```

6.在 vite.config.ts 中导入 svgIcon.ts,在 defineConfig()中加入 createSvg()，如下所示

```typescript
import { fileURLToPath, URL } from "node:url";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

//这里会报svgIcon.ts不在tsconfig.config.json文件列表中，在tsconfig.config.json的include里加
//"./src/icons/svgIcon.ts"就行，不加对npm run dev没影响
import { createSvg } from "./src/icons/svgIcon";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), createSvg("./src/icons/svg/")],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
});
```

7.配置完毕，我们可以用以下方式导入自己的 svg 图片

```html
<svgIcon name="你自己的svg文件名" :size="20"></svgIcon>
```

```html
<svg-icon name="password" class="svg-container" :size="15"></svg-icon>
```

还可以使用 main.ts 里设置的组件名，比如写的是 createApp(App).component('aaa', svgIcon)，你也可以这么导入

```html
<aaa name="你自己的svg文件名" :size="20"></aaa>
```

效果如下

![image-20221117145044507](https://c2.im5i.com/2023/01/10/YPNbW.png)

## 附录

<span id='import-type-node'></span>

### 导入@type/node

1.安装@type/node

> npm install @types/node

2.在 tscongfig.json 中加入"types": ["node"]，如下所示

```json
{
  "extends": "@vue/tsconfig/tsconfig.web.json",
  "include": ["env.d.ts", "src/**/*", "src/**/*.vue"],
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "@/*": ["./src/*"]
    },
    "types": ["node"]
  },

  "references": [
    {
      "path": "./tsconfig.config.json"
    }
  ]
}
```
