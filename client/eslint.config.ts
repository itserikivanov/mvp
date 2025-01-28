import js from '@eslint/js'
import pluginVitest from '@vitest/eslint-plugin'
import skipFormatting from '@vue/eslint-config-prettier/skip-formatting'
import { defineConfigWithVueTs, vueTsConfigs } from '@vue/eslint-config-typescript'
import pluginSimpleImportSort from 'eslint-plugin-simple-import-sort'
import pluginSortDestructureKeys from 'eslint-plugin-sort-destructure-keys'
import pluginVue from 'eslint-plugin-vue'
import ts from 'typescript-eslint'

// To allow more languages other than `ts` in `.vue` files, uncomment the following lines:
// import { configureVueProject } from '@vue/eslint-config-typescript'
// configureVueProject({ scriptLangs: ['ts', 'tsx'] })
// More info at https://github.com/vuejs/eslint-config-typescript/#advanced-setup

export default defineConfigWithVueTs(
  {
    name: 'app/files-to-lint',
    files: ['**/*.{ts,mts,tsx,vue}'],
  },

  {
    name: 'app/files-to-ignore',
    ignores: ['**/dist/**', '**/dist-ssr/**', '**/coverage/**'],
  },

  pluginVue.configs['flat/recommended'],
  ts.configs.recommended,
  js.configs.recommended,
  vueTsConfigs.recommended,

  {
    ...pluginVitest.configs.recommended,
    files: ['src/**/__tests__/*'],
  },

  skipFormatting,

  {
    plugins: {
      'simple-import-sort': pluginSimpleImportSort,
      'sort-destructure-keys': pluginSortDestructureKeys,
    },
    rules: {
      'simple-import-sort/imports': [
        'error',
        {
          groups: [['^\\u0000', '^@?\\w', '^[^.]', '^\\.']],
        },
      ],
      'simple-import-sort/exports': 'error',
      'sort-destructure-keys/sort-destructure-keys': 2,
    },
  },
)
