<script lang="ts" setup>
	import CustomPage from '@/components/custom-page/src/custom-page.vue'
	import Container from '@/components/subcontainer/src/container.vue'
	import defaultIcon from '@/assets/images/default_avatar.jpg';
	import TnAvatar from '@tuniao/tnui-vue3-uniapp/components/avatar/src/avatar.vue'
	import TnButton from '@tuniao/tnui-vue3-uniapp/components/button/src/button.vue'
	import { ref,reactive,ComponentInternalInstance,onUnmounted, getCurrentInstance } from "vue";
	import { useRoute,useRouter } from "vue-router";
	import QrcodeVue from 'qrcode.vue';
	
	import Countdown from '@/components/countdown/Countdown/Countdown.vue';
	import TnListItem from '@tuniao/tnui-vue3-uniapp/components/list/src/list-item.vue'
	import TnPopup from '@tuniao/tnui-vue3-uniapp/components/popup/src/popup.vue'
	import TnInput from '@tuniao/tnui-vue3-uniapp/components/input/src/input.vue'
	import vpay from './dtl'
	import { NOTIFY_URL,REDIRECT_URL } from '@/config'
	
	const Tlogo = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGQAAABkCAYAAABw4pVUAAAAAXNSR0IArs4c6QAAAAlwSFlzAAALEwAACxMBAJqcGAAAActpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IlhNUCBDb3JlIDUuNC4wIj4KICAgPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICAgICAgPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIKICAgICAgICAgICAgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIgogICAgICAgICAgICB4bWxuczp0aWZmPSJodHRwOi8vbnMuYWRvYmUuY29tL3RpZmYvMS4wLyI+CiAgICAgICAgIDx4bXA6Q3JlYXRvclRvb2w+QWRvYmUgSW1hZ2VSZWFkeTwveG1wOkNyZWF0b3JUb29sPgogICAgICAgICA8dGlmZjpPcmllbnRhdGlvbj4xPC90aWZmOk9yaWVudGF0aW9uPgogICAgICA8L3JkZjpEZXNjcmlwdGlvbj4KICAgPC9yZGY6UkRGPgo8L3g6eG1wbWV0YT4KKS7NPQAAHbZJREFUeAHtXXuQVNWZP7f73n5NzxMGhrdi0IhGSiUPUMKMSTYxiYkwM0iypYlJZAwMxGST3WRT2Qz5I2Xlj6QiDxeiq7VurYaeQRNTMVYSwTXR0kKjCKgRCKDIa4CZ6e7pd9/9/c7tM3TP9HsaGGQO3L6v8/jO9zrf+c537ggxnsYxMI6BcQyMY2AcA+MYeF9gQKtQL7Qus0sTaytU24VSzY+F2O2bK3HoW7YscaGAPQ5nCRgYrYSwvPnVbQ+5AieDj2nCnGxqIiJMm02IZAlgXIBZTU0Tsq/41cx93W2rV1SiF6MiSPuWLodvWVe0tXv9Mk9D7a9jg2GhkRbCxCFpNexMkIe/47P0pN7zGa+ZVF3W3Zn7dPDTy6l8PBf7PFtbql11VvWqe5zx32a3i5A/eNPW1lXbVmzaZGzu6IipnKWeWXNZqaury4ZDikFb9/rX7IZxTSIWCwMBelkVXpCFJPpidqfDnQhH/9Ld3rlImNARTJqmKFxSz8pG3osNDQZairT5Nvyz4XFeEwuGk+AWl8W9JcFwoWfWE5GoaXcaN7b5Nt7crWlPtW/Z4vAJES2nY9QvJSdKx1Nr1kRkQc280zqLi9nKiNsM8nbyLuLiqj174kOSIpFT/E9ZErJ98WISMtm+9f6bNLv2CY4dSPbim33f5bTFiQNNfKG1+/5ru9q++bfmxULfLkS81J6WJSHbW1qkNJjJxI90hwNtmhTPsuoqFeAxmt9ummbEWeO1a1ryB4Sxubk8M7NkJEI/UhLMpY/d9zG7w2iOBgcxsGscTy72ZET8QWDGbG3t/uWcLq0riclyyfgtucCJxj3SirAZtk4QBADIscOyLC5ukthAjKijusqmCf2bRMWTm6eUrMZLQqSysZf2bPiw3WZ7yUyUNo7bMJca0ymPoWqC8/K8Vt3CdBH2rpmMx+3mB59YsmZfc1eXvr2rq+ixpKxB3ZZMdhoet4gGg2hIK7qOhJkEE1mGcRGdU53MeiZpc9WR7x0rU+8Ve7Ae4DHnIKjysWyBBGKYMYe3yhCB0N3I+73GuXNL0kJFt3VGOtbPsptip7Dbasx4knOPgg2ykTgoETYTMMUKZs/sMwvnwnxmzjLvTEmMOJglChizISScTIhauyE8NnsxoCRgAtuTsfi7WtK42reso78UKSmOuwHnaV+9nJXbTO37RrWnJuoPRgE9Tay8ydBs4mQsIq6umSA6P/Ip4dIdIkE65lRfkl/z1lmZl6odEkETlN44EM/r9JTEcx2E+P2bL4tHD+0RUx0eEcOzzFzpJQSJEcNYMh2D/Lfw5ieXT5mibc/IkvumKIK0+7bY6F7+4qMbZoBdb4+H5LyjIKuzyxw3QpCMGsMpLp3QlBuSMf5mdt0kcXT/a2KGUwNB8gOL0UbD7B0+R9Hx5d9tXLf58x2naXHR8spfUoiiCHKisVEyhKEnb3NUe6ui/gDmHVpB6RjiInQARoCIJeKS28h1wzkxJ6CspAACcpYt8QWbGYI5VZaSY9h1CbsDzFXM0I6+wZ0SizlqqqbGAoGlqOrB7dulrh49Qegm6WppibdveaDB1ML/Gg/RY6KVZM6xE3aoLjtE3xo8bXlUVgoT6afhWEp/d7avU8xAuLMRLGfzoAqlxDRt/wav+CO+li5OnguyV0G1s3uuJUVJW7jTWe1tpH5ExSURhHBY/Ur1Lmcv3lcv9EQUUuL1zBH2CbezZyt2bCqokfIShNLB9Q6QmUPwV+Jh6U8skRjpSD6frJ4Oxzm7tmFGQvTRBBabr1+BmxFaMQOYvARR0tG+dcNduss5G+sdNEPylsmoffzGDiZOYr1kfmv3ulboaRMGUl43U07k0mdF6aBvH5Od71q45bRuPJWIgYTNLsfMf2E5WKtRap5cdeR8UX/6tHxnakc/j1n5HFAafpLSBvNcjV6oz8vkRnssFDGxorqgfevGxez7e1Ny+7iyDzLQc5s1a11Y02yc3DDRcVX2+EFLS04IYf5Ks5crncrmSB9a2Gv1nK2mvaN8cl6Te1LJAiMT2xuRVDvpL9QznpnQNmEmhyfRODlUvZLvi/uRTkfM3h1Q+d9BkWc3r+iIi46MXg7VlFVCmtd2ScS3b1m/RK9yfzwWCseBhby6b6jGYRfEJzxY0uw1EAxAZNL8tYEwNpjCQ2dep9/zvcqTese5jDI/hzWT91bWq+pX52F152qbcxAmwh6TrpU0DsnbasZLIxYMUUq+sLRn3Q0ghYmhICvus0oIHGKSEUybWCVZQs6GUov3Ge0UviExnEBCCNbyCX+fJIyUENptJSaMZaLaXSWcevG8wbb6ggHJ4WU0KZJJU7gdTtEfHhQ1YCRKShkJnTWjdoeDUtKJ8n9Vk+3hdY3AinIiQjoWYFr9PFzJJMeIfMMrKuYe/J2ZjbeF+oc8zMayhyKD4sGWdnFF00wgCqoEXJ4rkXiUpmA0LL779P+KI+GgcIPL0ZuyEpFQZlHVXso1b8Jhoc17fPmqN5q3denbWzJd85k9gsE89cgRjhXwxoj/gKlLhMl7VWu5Z3YmCp9WxgFnXsb98Pe8Zx4gn/mCuFYMWgpy6AyMoHwE9fBcsM1heVgmoRouFwHgK0hrHBNFaL/k961qFo+oLYMg7b61BkyyJGzm+Zrd9pnYYIgFMvKMqKHIB+RyJ0S+7AMGnleOQVaDrK/YZECSXKm21bkUOFgGi+XFNpczH8wRmwwI0cTyJY+tvxyxCXFqpPQC6cjWTjSqhXnta1yAgnTQTZKeJ71sydfksnIO6m3IhlxTKYdRZXkUVOdyYChFInMiBjrUTCbhmvc6bLp2R7Z8Q8jmqE991r7lvrlwG39DSsd5nJUTAZbeBiJ5lSJmtk4Uesa6WIc8UM+Zf4VKVvi9FdWo0UEL0nQs7fnFFIadchKuWhoiiHpg2uz3OGu9BihJx9VQRvW+0meFHIl0VE7FQPXggGXmsulYpdOhqgzhhWU1AYtb5azLe2C61tgdsp4q1OVGnU6oQB3tsD1FfEqQ9a/SvcyoT08m4ggZqp5oS+odfKMCR3gtzV7pYscC1Jee3DQxFokvR2gPoTQktMxV4UQk0AIiN9D+J1qICKoSDuBhrJsMJuOiH4P4IO775aBsir6wX0TxjqlYFULJesV/Ujw3OCDqQYw6tOfFmFKNccELwrhTY4tOOBAvQjhoiSmpLKUtCVhxP3a5yKdp3/jqQw/d+3DLnWE0K9EgCZKayidj4di3YAVUgyAxvM4YbIprJzOXQho5kR1m4tp1BIgOwAvaD+SeBuIjQBoRNBkS0ORwiZmeajHRWSUaXB7hxX2t0y08WHFkZP3UugmynoKSgjaZOGe59yM3i2AkJKKYC/VHwiIAU7gPJnQv5hYnIkFxBO+OxiPiFBfQ0PEGGA91IF4VCObCNaWViQxD+Nkvq3b5uPQfU8A1H40bVZ5pAS10FypYt2LzJn2z6Ihp1F9cnv3Sb341ORaLvgU3QW0ykaCvYYQ6K7ZlRQhaJyRGIBET76DT/0hERQNUxRWuKnGFt15cinX2adX1YmJVjajFhM8NpOtAABtmxxMIM4qCcKF4VHAZJhiNiOtmfEDUebxyTKGUFUqcr7xy6G1pPrsMhySQE8i2ox1aX4SV7pFQLCoCINKJwIA4PHBKHPCfEnsDp8UbIb84BELNgtq7xOEWtWAaJgY+kDSW0pOPSvvBBA9eYBvWlw4E3IkPPvXZNRFqKj3gPUopScRjkTsgHbXRQGmhPcOhoMgbKf18GB3cBg5scXnFrVNmi6sbp4vptROAfK/likBH/eGQ6A0OiHf6esVxqJWjg35xNBQQveBici2lKATOpEwfR117Pt8hCUJ9n88UJRxEFgNjfvbyn4Rv4LiYZbikjqY0TgJiJ+F+MpijCRLZVFUrJuNoqq4TlzdOE07DkET3g5GOgEBvnDgsnj12UDzi7xXX604xB+Uc6CcJU5gthmMJ93ASYtYeN9zuS6qCg1/Gk4de/EiDIeu65clNHmck+iqm9nOYCX3J6lLJUu2IR7Tvj0VDYh+2itzWOEt87gPzxJzGqXJWfcLfLw6ePi7ePnVUvIHznmCfeBV5ByANE6AWZoBz68GJXhyshxErRDrVE0caEnjd4lbM1GcUPVMPIeLlh3/8tTgMInOmTrVDziahImh3EATvw3EEUvwukAtEiY+CUFd7asWVCGy4vKEJwRmTxeSaehkX8N7pXvGn/bvEo+++hYHMFJeB2bjuztlzGYSJY/KtxyPRl7tbV80nMiXiXZH4XUa1dw5CezjvKJsYRN6bEPHPNM4Qv7x6gZg7ZZZ4r/+keGbvTrHt3bfFCxhcj6DzU4HwJnDZZKioLzg9Q5yukCXP6GwUHWWnyensbB+QxkGaSalFeZPnh74oP1ReMBmzBmpZ0jIoSHRacZMdmrgSsLMNDugc4w5DGl85/JY4eHCXqIdqW+htEIubZovmy64S9yy6RdzWf4N4fPeL4uEDu8RMqDLGDFijSx5gRr7iAlbCqHJf37Z1/fLupZ2P6dRbu4R5N3SZ1UsLrpFF8zxhR0iMf4CD//2qG8SSeQskCh9/7Xmxcud20YiOTwfXzXVXi3kAnPmJdE72KPKZKD4DAPOR76Qdhpt0FWW9kxny/yCjkjDWoxzxbFPObdi6BMCCgq0zXz0YZgJgvgr5OJ6dhqRt2P83sfLvL4l7L/mQWLnwZnlc23SJWLvjaVEN6SZMVi35QUp7iyLwy0D7ArDVeP6YbddVDbfqbucH5QIUjKG0zEVfOoDwf8B5d9fseWLpvIUATBN/2bdbLN3xlPhszUTxIYg/ASYCyH0kAv1LHAeY2BGWsf7JR1l/SuysrKNQrVbbZ9pnIXK69H+lYCXzUJKugHq6vXaS+M7+V8WDL/1J1r9g9pXi2x9aJA6AGWlJlg6jZmASzgF+IWKmW8AMtp+a0qhKwS9Ppf8cgfl61aTpQwWn1DSIhVV1Yj90Nwmhw8ZXYwKRwETgS+kAy3GwZiLSWG++g/kYjWhFGlpzDD7LlxRMCi5KJRFN+JmOY8wDzsSchslD1dAIUJI39LD4C1RsJjkHwiaTH2FeZruCkRFgTwVD8VWlchIpcyDeT+97XURklJAQlzVOEQ/ftFx8fOI0cRid4NhyFGd6b0kYzpbJdU7oZ3baQraFcIlsiXoL/UNEQDsGyjAxvJNmb65Dueaprtgm8ZmObNZ55p/VEcJBeDhZrEoZFoxJZijsXjAWVfJ0qN3nbrxVfPKKa61C+N2GfnvQJyaLbPKyuB9ZQDMRMkT4WnRsKViGfR5bkomkzjB61FKy2qK1QhPyDycOCcdzT4qvXd8irZI5k6aJtZ+8Tbx17B3x2nsHxOu9h8XugV6xC8QJoUwNOs/ZM60q6YVFpyRpiDwggoThIMtEuF3Iz4UiWkhxGAdcQVQokP1iRiQW4cKUAcTSbOU9yCfnRMxP4rAdtsb/aAzjhKVO+4CCPjBVH4gI+sBUdou5kPZrJkwV85pmSUOF8xmm/lBQ9Ox8XmxGiOmlME4oiSUnaEebzgEIEmGKZQRHtPdsaLPpug8TQggPelrCFgMFALmNVsshIKweAN9+2bXio7MuFzMbJqkscsw4PnBaWl4HTp8QhwZOikOYfB3E3INS1IumGYFOY7wGNr6b3AoiOXCmFcNZdA3qrgPxycHkfh1n2YmhVkgMGAxAKAnRB/P7KMY3Lr9ycCbiQyAoI/H9OHOOQwJNhORNhVdgBsaJmd46MaO6Qcysmyhm1jdK5kpfpTwR6Bc7MNl89O8vi9eD/eJKSE25xLA7dCMRw4TLFJ/Dtuo/c6aOzf/LoiSKptt9HE/oIkZfSnadkCh0CIaA2LdgNs4A4m6cOF1c1zRLXDl5uphW1ygnhGm4Q8ws3CggYh+4rQ9q4dRgQErBSVz3Q0X0YYLoxyw6ApOXM/4+XPfj7CfzoCJKERPP5HoSxzqgk3HRgAngRJjXNWBCFw4PXCn1QHw13DET4B2oBwEaMPPn7L8O99Vw16Qjn3WzhWP9p8TbJ94TO48dEi8cPyR2YA51OaSCM/cQpYkZS0tY0jUcUFUR1L+4p63zxZt/fx9CuZHUFxkUUSg7yTh6XSZROAhynKA1Rf3bi6roG7oGXDe3brKYg4nWdHBfo7dWukxUIEGu/iiOD8diMuiZ6ogry0MJvaCKk4qIPVKvcM2xhFLkBiGkWwZw5Ut0o/jBICfhPTiMOdS+U8fEbkxkd0OaD0PavNACk0BgeqAJF+dKEon5Kh3xDuvrTqcDAdn+pGY2b23tfIXEoPtkqC61lt7mW9cMK+K32Ahfja8TjMrJSB1NdUDVQuAHoJuPQTqPg7s5flwGDrsUJvE0+LKmwG0xES6MerhVql1u4QUHk1M9CDBwoPNEpjVmjOhdUQ/YPqPv6R+j34oz+AAcjX5KIaTzJCSTrpvDUEf7B/ulGX8MjFQPAkzBnKQWsEjfF4hN2WR95SVFjCgcBIlFjy//1i6lpVjfEEF4o6gkl3CF9n92Q3dDvzFqu+DWA5bPlRToyutLa4b6PAzkDKDTA0CUH9JEkWTIP6M7asDV1VAH3LnkxbjhBVLI21+/rllcMrFJmrN03edKlBgScBAOyQde+qN4DxwfwYDN6Bc/B210i6qPvrIw86KiKuSvhcqleuMaijWvODP2UCVbSjFXq4Wep4gRjfkB3wKoqd0K56pkBkH48KsPdbkevrMr3Nq98RrsuX4BsUQeuIqx+wT7QRRmVekyzuwUEztGwsiDd7gmEhX3cfAl0ThYUjtxjHga1tnuT90h5k6dVbQvi1EnX/vdf8HYCGL2bfGVgbaoUnUQgNYXJVh6jlPtsF1OBhWsFrRldDajiBnV3S4HVgsH0LNFPW337BxODGYfwWIkBkWop23lTiBoMQadAPUdoOMWtlEnEoL/mNhpRoIMQjqC4FQOjjHckwAkFE1hrvQ1QEIaYSC0YI2kHLXVCF/TNBwN0P11IIoHEkBioHlJcI51dDJyUYzXhItJwTr6bptRw+MhMfbG48nrSAyO2xwzZENpPyMIwnfS6pJEWb1DS2qLsPEkSIsAPaD6qmhiZ9XBiokKSgORQimhhHAVkedBHJSiUpOsQ9ZjzdqVBLAdVZuCYfTIHwFdBAEjDrhH9mrJ6EefWL5m38333eeU2zxGZM0iISoPiUKR8i1b9SoCihZAUoK62w2ZrzxRVJu5zulISr/OlT/XcyWZud6fhecRR5XHidCft+L22Id9y75zSqop9eGeLA1mlRCVjyJlWQCrXwcvzYfIvU3RA0uPEDVV5myfFUef7XYqUH8EbnUnlsP3aMnkoieWfLsvl5pKbysvQZgxTX29mbCZCyh6hpshjeePKOkdGKPXYagpJ74h9jbU1CLfsjUn1FyvELwFCcIKlPp6fGnnybi96sMQwb2kPl6dN0kp1LHz+J5fBXJFBwd3RQ3X9VRTxRKDMBdFEGa01FeX44kld/YZTmMBwuvfHCcKMZORovishjMyEPhzzPAs/O0Xv+4vhRisqWiCMLPa4vboLR29IMoi6Md9KaLILwkwz0WcIvh6gwPff/lDd1vnJ0kM7Lo1cllTufBUEkFYiTWmdDlIlIQ9Ph+SspsiilcXsfoyo87aamfMH/x9d2vnzcSTdEXNL/3rpCUThI1ZkkL19e2+qOFeEBnwP4cQIo4pFZ+nsL2xnfiNLK8j0u//na+t83OElfs+yv1UbFkEYaMkCkWSoomPCH8c6usZEIW+iYtDUhAtjSltGDG6juiAVFO3EC/N27aN2ITD58WmsgnCBjZDJNX+BojqJ2KBwaddEF28GlNEOQsTQrp74676ahcG8N9gYUmqKWtHVAtXXctOoyIIW6VoEhBe+6A/w/2BpzC4gSjnfkZPGIYnhpIy0oVLAWccJcNzlXTPuWkS2sAI9/X/NwbwW1k62/a0kmpNZR41QVgP95VISUGfu9tWfRbhqM9Qr0qiSNEuB7TKlKFPjL4srstUIIEYZgLEsEcHgxu7W1d/hXU28zN+w/YKlttWRQjCxi1J2SYlheor4g884aqroUOSyxzkqvOSSAceFQCAAhZ31dXqUM0/BzGwQ9mUwerbS/imYiEkVIwgbIh75pT66mlbvSTcF/gf2uZ4xXCMCuCErZyXRPgT+P4V1JT/F7Cm5Gcy2n0++WG3SkJUUYIQMIouRZjXUF+3I154A0Uct4xHvhCJIpkJcy09OhDYiD7xawyIQ2iX2zh4XclUcYIQOIowvMT4YqSpYdDrjAVCv8TEiUQ6r+qrDMTJLeH8YnWkP/DT7nalpkgMn3xXRp15i5wVgrBFzOgTFGnIhOZrW3VPtN//c3SMA30CYlK2pFRmbM6LE/WSCLdju4AtPBD8AUzbH5LBusRa7WwRgw2fNYKwcoso1jc9qHfDA/71+CqdDqSWTRRFSXVmO2chKWKIWDi0uqdt1b2U+K61a7ViPmQ5GnjOKkEImCRKSn1hoF8d8Q/+gvoYROF6bFGxlwr50lpKLeFyKZehONbBdXgrmv7M2SpVRriOJAb2/4lYJNLR07ZmvXQStrcnu1J/wGY0CC9U9qwTRBGFok71xUERroZ7YX1Z0dLWYJ8XTjXT5pzCiaV9JivozYoYkV/yQRQJo0esa0SzIIiByY2/3qDKywf5f/DVI6gpxHREg6EVPa2dm2k1yk/znaP5lLSG8sNYmbcUdfntWvPHwqdpP2jvXn8K+vln2JfC8A9+WTknc8gZNj6+UodYqdePHBRhxFQxipEEyJrwmBLECHkZnM2JITCtJC1HGbmVT3e5RDwcvgsxUw9Iv1RzM6Ou8xbNWl+ZD3P0qMzaiimGgbF57Vo7LTHEft3tqHLdj2Vh4Fyqr7wfKiDFjiDi8Dh283KrcvYIFKnYpFQQw1XINwtRkPmlxIxrNrtOyYiHQ9/wta5+ULrP+aGxAnQspsul5Dn3BCF0JMp2EAVzFkRJ3o2dqPenvpZdcDsEox/lNoIie0nWpuskT5LbAWTkv6l9GQP4o5bHVv7RmnMmGQq+c6ayVIPyDBWw3TQTzduE3tOy+j/xZ/fsDrdrPb5cR3jkoJqRP+2G44gV9572sPzLGP7knYHAchOC9GnshP2jXHJtbjlv86XzQxAiMEWUFZumGJvbOja0+dZHdI/zVwg1YnRnzq3ZFWRZKRncfoEwnRZf25pnZRxalmjC8uldesnzo7Iy4dSu37RJf5lfxelZ93W7y/0AYokhJ0l+877kPSqZVWe5s4aYKD9KiZ3HILz5KczAt5cajJCl5oo8yjuIVqSFIio58uST5oqpU41Hbl+548rWf9qJvR/Y0WXXcabqqByMFvtFsFvMiX2V0HyKGNy0tGpMLD+PBQlRJKMrG1Eay6Lco4Iv2m3jx2bK3TikKh12xkYZfIgyEk1go8zira2r/5otAn1YmXN6m9P2P6dQWI2ZJAZDV6lCMMg0gxhJqBaqLUrKaFMU8x4Sox+q8AYSg2oqWwT6aBsaTfmxJCFD/SBRSByYxJ/GfPEpbEvTpKRgp1z+2d1QFcMvIphjOPFNkaP4JvLiXy/t/LtqY3jG830/JglCpChVwq8bYPPmM5b64m6uwn9IZhhSI9go48Rf4uyNxbXrfvOlVe+ouoflGxO3Y0llZSBERd7zT2Kn1Fc8tXGoBPUldy05YUof0UxtAYlByRhraiq942OWIASSakt6WpeufBZ6/0bo/1N2l/xrlsVYRHKjDDwAB7WkDXszVu21QjuXFVM2HUfn9HrMqqx0LCgV075l00xTi72ke1yTgWjGfjEGLFsKI+bYhTDXg0nNWLC1teOIqiNb5rH0bExLiEIUVQwRir8JeMi02RZiPDiGLxhZsV8jWAr7+ao8JMab2I73MRJjrKsp1U+eLwiCEFBrTOly9LSu3K/r9o9ZRHFZ+x6HiII4W68X+/kGX3IPhuZDTR29ENQU+6fSBUMQAsx4Ym6YfGzJNw9ghWKBRRTse8TnS/A6zOA8bAd4QUv0Lnrkju8FOc9guKvq7IVwHuKtCwFYBaMaD9q3rLvUtGnPY7xoolmMlcgXNLO3WUXn86zKjJ/PMgYoKWyitWfjbLjv97T1rP+zekY1dZabH68+GwZUlCSCD4ZUr3qWLf/4s3OAgeZUlCSbgjVVOc/wOYD9fduEDJ5Ik5L3bUfHOzaOgYseA/8PSlxJyHnSx6sAAAAASUVORK5CYII=";
	
	const qrcodevalue = ref("TUe53VJrJqp9fWgRTGaPUtwanw1ZfnM2qM")
	const size =ref(150)
	const currentCollapse = ref(false)
	const winningCollapse = ref(false)
	const {proxy} = getCurrentInstance() as ComponentInternalInstance;	
	const that:any = proxy;
	that.$vcommon.checkExpire(); 
	const route = useRoute();
	const router=useRouter()
	const QueryId = route.query.id	
	let memberinfo = that.$vcache.vget("memberinfo")
	const tradeInfo = ref()
	const currentTime = ref(Math.ceil(Date.now()/1000));
	const restTime = ref(10*1000*3600);//十分钟倒计时
	const endCountTime = ref(false);	
	const showNumberReword = ref(false)
	const rewordNumer = ref(0)	
	const finishMaskTxt = ref("订单超时,<br>点击重新获取")
	
	const dtlData = reactive({data:[]})	
	const rewordsData = ref()
	const rewordNumList =ref([]);//付款号
    const userIsReword = ref(false);//用户是否中奖
	if(!QueryId){
		uni.showToast({
		  title: '无效参数',
		  icon: 'none',
		})
		router.back();
	}
	const showNormalPopupWallet = ref(false)	
	const walletAddressInput = ref()
	const openNormalPopupWallet = () => {
		if(memberinfo.wallet_address) walletAddressInput.value = memberinfo.wallet_address
		showNormalPopupWallet.value = true
	}	
	
	const listgetdtl =  "/g/getdtl"
	const initData =async () => {
		that.$vcache.vdelete("tradeInfo");
		let resData = []
		resData =await that.$request.post({url:listgetdtl,data:{
			tid:parseInt(QueryId),
			cuid:memberinfo.id,
		}});
		dtlData.data = resData.itemdata;
		rewordsData.value = resData.rewordorder
		 let findrewordmem = rewordsData.value.find(it=>it.cuid == memberinfo.id)
		 if(findrewordmem) {
			 userIsReword.value = true
			 //rewordNumList
			 //is_limit_N			 
			 if(resData.is_limit_N == 2){				 
				 //pay_count
				 rewordsData.value.forEach((items,indexs)=>{
					
					 if(indexs < resData.itemdata.pay_count){
						 let s = {
							 pay_number:items.payed_number
						 }
						 rewordNumList.value.push(s)
					 }
				 })
			 }
			 
		 }
		 
		 
		
		
		// if(resData.reword_count>0){
		// 	for (let it = 1;it<=resData.reword_count;it++){				
		// 		let pstring = "第"+it+"次中奖"+resData['reword'+it]
		// 		awardLists.value.push(pstring)
		// 	}
		// }
		
	}
	initData()
	//检测是否过期
	const checkExpireTime = async(tradeId = "") =>  {
		if(tradeId){
			let expireUrl = "/pay/check-status/"+tradeId;
			let resData =await that.$request.get({url:expireUrl});
			if(resData.status ==1){return true;}
			if(resData.status ==2){
				uni.showToast({
				  title: '支付成功',
				  icon: 'none',
				});
				return false;
			}
			if(resData.status ==3){return false;}
		}
		return false;
	}	
	const startBtnGame =async () => {
		if(!memberinfo.wallet_address){
			openNormalPopupWallet()
		}else{
			 tradeInfo.value =await that.$vcache.vget("tradeInfo")
			//检测是否过期
			if(tradeInfo.value){				
				let restLeftTimeLeft = that.$vcache.vget("tradeCreateTimeLeft")
				let ctime = Math.ceil(Date.now()/1000)
				let cbets = tradeInfo.value.expiration_time - ctime-restLeftTimeLeft;				
				let getExpire = await checkExpireTime(tradeInfo.value.trade_id);
				if(!getExpire || cbets>=0){
					that.$vcache.vdelete("tradeInfo")
					currentCollapse.value =  false
					tradeInfo.value = ""
					endCountTime.value = true;
					restTime.value=0
				}
				
			}else{
				//创建支付订单
				tradeInfo.value = await vpay.setReqpost(QueryId,parseFloat(dtlData.data.pay_unit),NOTIFY_URL,REDIRECT_URL) || ""
				qrcodevalue.value = tradeInfo.value.token
				currentTime.value = tradeInfo.value.created_at;				
				that.$vcache.vset("tradeInfo",tradeInfo.value)				
				currentCollapse.value = true
				that.$vcache.vset("tradeCreateTimeLeft",Math.ceil(Date.now()/1000)-tradeInfo.value.created_at)	
				let cbet = tradeInfo.value.expiration_time -currentTime.value ;
				restTime.value = cbet *1000;
			}
			
			
		}
		
	}
	//重新获取qr
	const refreshQrcode =async () =>{
			tradeInfo.value = await vpay.setReqpost(QueryId,parseFloat(dtlData.data.pay_unit),NOTIFY_URL,REDIRECT_URL) || ""
			qrcodevalue.value = tradeInfo.value.token
			currentTime.value = tradeInfo.value.created_at;				
			that.$vcache.vset("tradeCreateTimeLeft",Math.ceil(Date.now()/1000)-tradeInfo.value.created_at)
			let cbet = tradeInfo.value.expiration_time -currentTime.value ;			
			restTime.value = cbet *1000;			
			that.$vcache.vset("tradeInfo",tradeInfo.value)	
	}
	
	const onFinish = () => {
		//currentCollapse.value =  false
		tradeInfo.value = ""
		endCountTime.value = true;	
		restTime.value=0
	}
	
	const WincontrollerBtn=()=>{		
		winningCollapse.value = !winningCollapse.value
	}
	
	const walletSubmit =async()=>{
		if(!walletAddressInput){
			uni.showToast({
			  title: '请输入钱包地址',
			  icon: 'none',
			})
			return ;
		}
		const walletUrl = "/g/editwalletaddress"
		let resData =await that.$request.post({url:walletUrl,data:{
			wallet_address:walletAddressInput.value,
			cuid:memberinfo.id
		}});
		if(resData.id >0){
			showNormalPopupWallet.value = false
			that.$vcache.vset("memberinfo",resData)
			uni.showToast({
			  title: '添加成功',
			  icon: 'none',
			})
			return ;
		}
		
	}
	
	const Ispayed = ref(false);
	
	const payedInfo = ref([]);
	const payedMember = ref([])
	//每分钟循环
	const timeStart = ref(true)
	const timesRef = ref();
	const timesFunc = async()=>{
		if(QueryId){
			let getRestTimes = "/pay/getpayedinfo/"+QueryId
			timesRef.value = setInterval(async () => {	
				//测试
				//await vpay.setReqpost(QueryId,100,NOTIFY_URL,REDIRECT_URL) || ""
				
				if(timeStart.value){
					let resData =await that.$request.get({url:getRestTimes,showLoading:false});
					//console.log(resData,"resData")
					
					if(resData.item_data.length >0){
						Ispayed.value = true
						payedMember.value= resData.member_data						
						let s = [];
						//console.log(resData.item_data,"resData.item_data")
						resData.item_data.forEach(it=>{
							let p = it;
							p.member = payedMember.value.find(item=>item.id == it.cuid)
							s.push(p)
							
							if(it.is_reword == 1 && resData.item_data.trade_id == tradeInfo.value.trade_id){
								showNumberReword.value=true
								rewordNumer.value = it.payed_number
							}
						})
						payedInfo.value= s
					}else{
						Ispayed.value = false
					}
					//console.log("payedInfo",dtlData.data)
					
					 if(resData.status == 2){							
						 timeStart.value = false;
						 finishMaskTxt.value = "该订单次数已用完"
						 return ;
					 } 
					 if(resData.status == 1 && resData.rest_times>0){						
						 Ispayed.value = true						 
					 }
					 if(resData.item_sta == 2){
						  finishMaskTxt.value = "订单超时,<br>点击重新获取"
					 }
					console.log(" resData.item_data", resData.item_data)
					
				}else{
					clearInterval(timesRef.value)
				}
			}, 2000);
		}else{
			clearInterval(timesRef.value)
		}
		
	} 
timesFunc()
onUnmounted(() => {
     clearInterval(timesRef.value)     
});
</script>
<template>
	<!-- 基础弹框 -->
	<TnPopup v-model="showNormalPopupWallet"
	close-btn
	close-btn-position="right-top"
	open-direction="center"
	>
	  <view class="popup-content center"> 
	   <Container title="请输入您的钱包地址">
	     <view class="input-container">	      
	       <view class="input-item">
	         <TnInput
	           v-model="walletAddressInput"
	           placeholder="请输入钱包地址"
	           focus
	         />
	       </view>
	     </view>
		  <view class="sub-wrap disflex spaceBetween">
			  <view></view>
			 <TnButton
			   size="lg"
			   bg-color="gradient-bg__cool-6"
			   text-color="tn-white"
			   @click="walletSubmit"
			 >
			   提交
			 </TnButton>
		 </view>
	   </Container>
	   </view>
	</TnPopup>
	

<CustomPage title="详细信息" page-bg-color="tn-gray-light" >	
	<view class="page-container">		
		<Container title="基础信息">
			<view class="">
				<view class="font14 disflex width100 spaceBetween paddding20 tn-white_bg borderRadius10">
					<view class="disflex justaligncenter">
						<TnAvatar  :url="defaultIcon" />
						<view class="ml10">{{dtlData.data.cuname || ""}}</view>
					</view>					
				</view>
				<view class="mt10 tn-white_bg width100 paddding20 borderRadius10">
					<view>游戏名称：{{dtlData.data.ctitle|| ""}}</view>
					<view class="mt6">创建时间：{{dtlData.data.created_at|| ""}}</view>
					<view class="mt6">发起人：{{dtlData.data.cuname|| ""}}</view>					
					<view class="mt6">总筹额度：{{dtlData.data.total_U|| ""}}U</view>
					<view class="mt6">已筹额度：{{dtlData.data.charged_U >dtlData.data.total_U ?dtlData.data.total_U :dtlData.data.charged_U || ""}}U</view>					
					<view class="mt6">每注额度：{{dtlData.data.pay_unit|| ""}}U</view>
					<view class="mt6">该游戏总中奖设置：总{{dtlData.data.reword_count|| ""}}次中奖</view>
					<view class="mt6">公布预中奖号：{{dtlData.data.reword_all||""}}</view>
					<view class="mt10 mb10">游戏说明:{{dtlData.data.breif|| ""}}</view>
				</view>				
				<view class="mt10 tn-white_bg width100 paddding20 borderRadius10">
					<view v-if="userIsReword">您的付款号：
						<text v-if="rewordNumList.length>0">
						   <text v-for="(itemss,idx) in rewordNumList" :key="idx">
							   {{itemss.pay_number}},
						   </text>
						</text>
					</view>
					<view class="mt6">是否中奖：
					<text v-if="userIsReword" class="tn-gradient-bg__cool-4" style="padding:5px 9px;border-radius:2px;color:#fff">
						中奖
					</text>
					<text v-else  class="">未中奖</text>
					 					
					</view>					
				</view>				
				<view class="mt10 tn-white_bg width100 paddding20 borderRadius10">
					<view v-if="dtlData.data.tg_kefu">客服：{{dtlData.data.tg_kefu|| ""}}</view>
					<view v-if="dtlData.data.tg_group" class="mt6">点击入TG群：{{dtlData.data.tg_group|| ""}}</view>					
				</view>	
				
				<view class="mt10 tn-white_bg width100 paddding20 borderRadius10">
					<view>我的钱包地址:</view>
					<view class="mt10 mb10 textcenter">{{memberinfo.wallet_address}}</view>
					<view class="disflex spaceBetween mb10">
						<view></view>
						<view class="mr20">
							<TnButton  type="success" debounce @click="openNormalPopupWallet">
							   修改地址
							</TnButton>	
						</view>
					</view>
				</view>	
						
							
				<view class="mt20 mb20 width100 disflex spaceBetween ">
					<view></view>
					<view class="mr20">
						<TnButton debounce @click="startBtnGame">
						   参与游戏
						</TnButton>	
					</view>
				</view>
				
				<view class="list-container mt20 mb20">
				  <view  class="list-item">
				    <TnListItem @click="()=>WincontrollerBtn()" right-icon="right">中奖信息</TnListItem>
				  </view>
				</view>
				
				<view v-if="winningCollapse" class="mt10 mb20">
					<Container title="中奖用户信息">
					  <view class="collapse-container">
						<view :class='[winningCollapse?"collapse-item":"","tn-white_bg"]'>
							<view v-if="winningCollapse" :class='["textcenter mt20 mb20", currentCollapse?"item-content":""]'>
								
								
								<view v-if="Ispayed" class="mt10 mb10">	
									<view v-for="(item,index) in payedInfo " :key="index">
										<template v-if="dtlData.data.reword_count > index">	
											参与人:{{item.member.cuname}}&nbsp;&nbsp;
												<text v-if="dtlData.data.sta == 2">												 
													 付款号:{{item.payed_number}}&nbsp;&nbsp;
													<text v-if="item.is_reword == 1" class="redColor">中奖&nbsp;</text>																					  
												</text>
											付款金额:{{Math.floor(item.actual_amount)}}U
										</template>										
									</view>
								</view>
								
							</view>				    
						</view>
					  </view>
					</Container>
				</view>			
				
				
				
				
				<view v-if="currentCollapse">
					<Container title="订单支付">
					  <view class="collapse-container">
						  
						<view :class='[currentCollapse?"collapse-item":"","tn-white_bg"]'>
							<view v-if="currentCollapse" :class='["textcenter mt20 mb20", currentCollapse?"item-content":""]'>
								 <view class="paylogo">
								      <img :src="Tlogo"  class="img-box"/>
								  </view>
								  
								  <view class="mainTxtColor  mb10 mt10">
									  订单编号: {{tradeInfo.trade_id || ""}}
								  </view>
								  <view class="secondTxtColor  mb10 mt10">
									  当前USDT支付区块网络协议为TRC20									  
								  </view>
								  <view class="redColor mb10 mt20">
									  ！!到账金额需要与下方显示的金额一致，否则系統无法确认！！
									  
								  </view>
								  <view class="redColor mb30 mt10">
									  尝试点击钱包地址或金额可直接复制👇									  
								  </view>
								  <view class="redColor font22 mb30 mt10">
									您需要支付{{tradeInfo.actual_amount || 0}} USDT
								  </view>
								  <view >
									  <view  class="qrcodeBoxWrap">
										  <view class="">
											   <qrcode-vue :value="qrcodevalue" :size="size" level="H" />											   
										  </view>
									  </view>
									 						  
										<view v-if="endCountTime" class="endTime-wrap">
											<view @click="refreshQrcode" class="endTimeBox">
												<view class="font16 fontBold" v-html="finishMaskTxt"></view> 
											</view>
										</view>
								  </view>
								
								<Countdown
								      title="有效期10分钟"
								      :value="restTime"
								      format="HH:mm:ss"
								      finishedText="Finished"
								      @finish="onFinish">
								      <template #prefix>剩余时间</template>
								      <template #suffix></template>
								    </CountDown>
								
							</view>				    
						</view>
					  </view>
					</Container>
				</view>
				
				
			</view>
		</Container>
	</view>

  <TnPopup v-model="showNumberReword" 
  close-btn
	close-btn-position="right-top"
	open-direction="center">
    <view v-if="rewordNumer>0" class="popup-content center font30 rewordfontBox"> 您的号码为：
	 {{rewordNumer}}
	 </view>
  </TnPopup>
  
</CustomPage>	


</template>
<style lang="scss">
@import "./dtl.scss";
.rewordfontBox{
	padding:50px auto;
	color:#ff71d2;
}
</style>
