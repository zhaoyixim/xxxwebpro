import { withInstall, withNoopInstall } from '@tuniao/tnui-vue3-uniapp/utils'
import Container from './src/demo-container.vue'
import subContainer from './src/sub-demo-container.vue'

export const TnContainer = withInstall(Container, {
  subContainer,
})
export default TnContainer

export const TnSubContainer = withNoopInstall(subContainer)

export * from './src/container'
export * from './src/subcontainer'
