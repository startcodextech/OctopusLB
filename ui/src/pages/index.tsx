import React, {FC} from "react"
import {graphql, HeadFC, PageProps} from "gatsby"
import {LayoutDashboard} from "@layouts/dashboard";

const IndexPage: FC<PageProps> = () => {
  return (
    <>
      <LayoutDashboard>
          hola
      </LayoutDashboard>
    </>
  )
}

export default IndexPage

export const Head: HeadFC = () => <title>ManagerLB</title>

export const query = graphql`
  query ($language: String!) {
    site {
        siteMetadata {
            title
        }
    }
    locales: allLocale(
      filter: { ns: { in: ["common", "login"] }, language: { eq: $language } }
    ) {
      edges {
        node {
          ns
          data
          language
        }
      }
    }
  }
`;