package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)


func ChatLoop(session ChatSession) {
	fmt.Printf("🤖 Chatting with %s — session: '%s' (type 'exit' to quit)\n", session.Model, session.Name)
	reader := bufio.NewReader(os.Stdin)
	history := session.Messages

	for {
		fmt.Print("You: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		// 🔍 Handle search command
		if strings.HasPrefix(input, "search:") {
			query := strings.TrimSpace(strings.TrimPrefix(input, "search:"))
			fmt.Println("🌐 Searching for:", query)

			results, err := SearchWeb(query)
			if err != nil {
				fmt.Println("❌", err)
				continue
			}

			fmt.Println("🔍 Top Results:\n", results)

			input = "Here's some information I found:\n" + results
			history = append(history, Message{Role: "user", Content: input})

			reply, err := SendChatRequest(session.Model, history)
			if err != nil {
				fmt.Println("❌ Error:", err)
				continue
			}

			fmt.Println("AI:", reply)
			history = append(history, Message{Role: "assistant", Content: reply})
			continue
		}

		// 🖼️ Handle image command
		if strings.HasPrefix(input, "img:") {
			imgPath := strings.TrimSpace(strings.TrimPrefix(input, "img:"))
			base64Image, err := EncodeImageToBase64(imgPath)
			if err != nil {
				fmt.Println("❌ Failed to read image:", err)
				continue
			}

			images := []string{base64Image}
			fmt.Print("📝 Enter your question about the image: ")
			imgPrompt, _ := reader.ReadString('\n')
			imgPrompt = strings.TrimSpace(imgPrompt)
			prompt := Message{
				Role: "user",
				Content: imgPrompt,
				Images: images,
			}

			history = append(history, prompt)
			reply, err := SendChatRequestWithImages(session.Model, history)
			if err != nil {
				fmt.Println("❌ Error:", err)
				continue
			}

			history = append(history, Message{Role: "assistant", Content: reply})
			continue
		}

		// 💬 Default: normal text input
		history = append(history, Message{Role: "user", Content: input})
		reply, err := SendChatRequest(session.Model, history)
		if err != nil {
			fmt.Println("❌ Error:", err)
			continue
		}

		fmt.Println("AI:", reply)
		history = append(history, Message{Role: "assistant", Content: reply})
	}

	session.Messages = history
	if session.Timestamp.IsZero() {
		session.Timestamp = time.Now()
	}

	SaveChat(session)
	fmt.Println("💾 Chat saved.")
}







