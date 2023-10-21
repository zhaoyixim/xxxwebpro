import { buildProps } from '@tuniao/tnui-vue3-uniapp/utils'

import type { ExtractPropTypes } from 'vue'
import type Container from './container.vue'

export const ContainerProps = buildProps({
  /**
   * @description 演示内容标题
   */
  title: {
    type: String,
    required: true,
  },
  /**
   * @description 标题带边距
   */
  titlePadding: Boolean,
})

export type ContainerProps = ExtractPropTypes<typeof ContainerProps>

export type TnContainerInstance = InstanceType<typeof Container>
