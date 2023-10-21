import { buildProps } from '@tuniao/tnui-vue3-uniapp/utils'

import type { ExtractPropTypes } from 'vue'
import type SubContainer from './subcontainer.vue'

export const subContainerProps = buildProps({
  /**
   * @description 演示子内容标题
   */
  title: {
    type: String,
  },
})

export type SubContainerProps = ExtractPropTypes<
  typeof subContainerProps
>

export type TnSubContainerInstance = InstanceType<typeof SubContainer>
