import "@testing-library/jest-dom";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import Page from "@/app/(auth)/signup/page";

jest.mock("@/lib/helpers/passwordValidator", () => ({
  __esModule: true,
  default: jest.fn(() => false),
}));

describe("Page component", () => {
  test("renders correctly", () => {
    const { getByText, getByPlaceholderText } = render(<Page />);

    // Assert that necessary elements are present
    expect(getByText("Sign up")).toBeInTheDocument();
    expect(getByPlaceholderText("abc123")).toBeInTheDocument();
    expect(getByPlaceholderText("abc123@nyu.edu")).toBeInTheDocument();
    expect(getByPlaceholderText("Enter password")).toBeInTheDocument();
    expect(getByPlaceholderText("Re-enter password")).toBeInTheDocument();
    expect(getByText("Already have an account?")).toBeInTheDocument();
    expect(getByText("Log in")).toBeInTheDocument();
  });

  test("validates password correctly", async () => {
    const { getByTestId, getByPlaceholderText } = render(<Page />);

    // Enter invalid password
    fireEvent.change(getByPlaceholderText("Enter password"), {
      target: { value: "weakpassword" },
    });
    fireEvent.change(getByPlaceholderText("Re-enter password"), {
      target: { value: "weakpassword" },
    });
    fireEvent.click(getByTestId("submitButton"));

    // Assert that passwords is invalid error is displayed
    await waitFor(() =>
      expect(getByTestId("errorMessage")).toBeInTheDocument()
    );
  });

  test("validates re-entered password correctly", async () => {
    const { getByTestId, getByPlaceholderText } = render(<Page />);

    // Enter valid password but different re-entered password
    fireEvent.change(getByPlaceholderText("Enter password"), {
      target: { value: "StrongPassword123!" },
    });
    fireEvent.change(getByPlaceholderText("Re-enter password"), {
      target: { value: "DifferentPassword123!" },
    });
    fireEvent.click(getByTestId("submitButton"));

    // Assert that passwords do not match error is displayed
    await waitFor(() =>
      expect(getByTestId("errorMessage")).toBeInTheDocument()
    );
  });
});
