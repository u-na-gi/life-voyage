/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "新規ブログテンプレートを作成します。",
	Long: `新規ブログテンプレートを作成します。
	あらかじめ環境変数にsetしといてほしいです。

	fishをお使いの場合:
	set -Ux LIFE_VOYAGE_PATH <Full Repository Path>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// フォルダ名とslugは共通のuuidを使う
		//

		slug := uuid.New().String()
		// 現在の日時をローカルタイムゾーンで取得
		localTime := time.Now().Local()
		// ISO 8601形式で出力
		date := localTime.Format(time.RFC3339)

		// 環境変数でレポジトリのパスは入れてもらってるものとする。
		repo := os.Getenv("LIFE_VOYAGE_PATH")

		if repo == "" {
			fmt.Println(`Please set LIFE_VOYAGE_PATH.
			
			c.e:
			set -Ux LIFE_VOYAGE_PATH <Full Repository Path>
			`)
			return
		}

		// post配下にファイルを作成する
		postPath := filepath.Join(repo, "content", "post")

		if err := os.Mkdir(filepath.Join(postPath, slug), 0755); err != nil {
			log.Fatal(err)
		}

		blogTemplate := fmt.Sprintf(`---
title: 
description: 
slug: %s
date: %s
image: 
categories:
    - 
weight: 1      
---`, slug, date)

		contentPath := filepath.Join(postPath, slug, "index.md")

		f, err := os.Create(contentPath)
		if err != nil {
			log.Fatal()
		}
		defer f.Close()

		_, err = f.WriteString(blogTemplate)
		if err != nil {
			log.Fatal(err)
		}

		msg := fmt.Sprintf(`created!
code -n %s
code --reuse-window %s`, repo, contentPath)

		fmt.Println(msg)

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
