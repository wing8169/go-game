// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.432
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/wing8169/go-game/views/layout"

func renderScreen() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_renderScreen_37bf`,
		Function: `function __templ_renderScreen_37bf(){const element = document.getElementById("me");
    let start, previousTimeStamp;
    let done = false;
    let keyPressed = undefined; // w / a / s / d
    const speed = {
        "w": 5, "a": 5, "s": 5, "d": 5,
    };

    function getKeyPress(){
        document.addEventListener("keydown", (event) => {
            const key = event.key
            if(["w", "a", "s", "d"].includes(key)){
                keyPressed = key;
            } else {
                keyPressed = undefined;
            }
        });
    }

    function step(timeStamp) {
        if(done) return;
        if(!!keyPressed) {
             console.log("halo23")
            element.style.transform = ` + "`" + `translateX(${speed[keyPressed]}px)` + "`" + `;
        }
        console.log("halo")
        window.requestAnimationFrame(step);
        
        // if (start === undefined) {
        //     start = timeStamp;
        // }
        // const elapsed = timeStamp - start;

        // if (previousTimeStamp !== timeStamp) {
        //     // Math.min() is used here to make sure the element stops at exactly 200px
        //     const count = Math.min(0.1 * elapsed, 200);
        //     element.style.transform = ` + "`" + `translateX(${count}px)` + "`" + `;
        //     if (count === 200) done = true;
        // }

        // if (elapsed < 2000) {
        //     // Stop the animation after 2 seconds
        //     previousTimeStamp = timeStamp;
        //     if (!done) {
        //         window.requestAnimationFrame(step);
        //     }
        // }
    }

    window.requestAnimationFrame(step);
    getKeyPress();}`,
		Call: templ.SafeScript(`__templ_renderScreen_37bf`),
	}
}

func Index() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, renderScreen())
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body onload=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 templ.ComponentScript = renderScreen()
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3.Call)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"bg-black flex flex-col items-center justify-center w-screen h-screen\"><div class=\"rounded-full w-[100px] h-[100px] bg-red-500\" id=\"me\"></div></div></body>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.MainLayout().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
