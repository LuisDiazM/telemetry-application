from app.dependency_injection import AppContainer
import sys


if __name__ == "__main__":
    container = AppContainer()
    application = container.app()
    try:
        application.start()
    except KeyboardInterrupt:
        print('Interrupted')
        sys.exit(0)