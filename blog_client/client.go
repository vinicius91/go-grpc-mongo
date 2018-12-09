package main

import (
	"context"
	"fmt"
	"github.com/vinicius91/go-grpc-mongo/blogpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Starting Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("An error ocurried while tring to connect %v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	blog := &blogpb.Blog{
		AuthorId: "Vinicius",
		Title: "Title of First Blog",
		Content: "Content of First Blog",
	}
	writeBlog(c, blog)
	readBlog(c, "dasdsad")
	readBlog(c, "5c0d71555192f6ca2002cee5")


}

func writeBlog(c blogpb.BlogServiceClient, blog *blogpb.Blog)  {

	fmt.Println("Creating the blog")
	createRes, err := c.CreateBlog(context.Background(), &blogpb.CreteBlogRequest{
		Blog: blog,
	})
	if err != nil {
		fmt.Printf("Error while creating the blog: %v", err)
	}

	fmt.Printf("Blog created: %v", createRes.GetBlog())
}

func readBlog(c blogpb.BlogServiceClient, blogID string)  {

	fmt.Println("Reading the blog")
	res, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: blogID,
	})
	if err != nil {
		fmt.Printf("Error while reading the blog: %v", err)
	}
	fmt.Printf("Blog Found: %v\n", res.GetBlog())
}